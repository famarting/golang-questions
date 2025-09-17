package main

import (
	"context"
	"fmt"
	"time"
)

// This started code is only a guide, you are free
// to use it, modify it or delete it. Do not feel
// the need to align with this structure in anyway.
// Approach the problem as you see fit.

func main() {
	fmt.Println("running...")

	// Below is some simple code to exercise your PubSub broker.
	// You can modify this to demonstrate more advanced features
	// as your implementation progresses.
	broker := NewBroker()

	topic := "mytopic"
	sub1 := broker.Subscribe(topic)

	numMessages := 10
	done := make(chan struct{})

	go func() {
		received := 0
		for {
			if received >= numMessages {
				done <- struct{}{}
				fmt.Printf("sub recevied all %d messages\n", numMessages)
				return
			}
			m, ok := sub1.Receive(context.TODO())
			if ok {
				fmt.Printf("sub received message: %s\n", m)
				received++
			}
		}
	}()

	for i := 0; i < numMessages; i++ {
		broker.Publish(context.TODO(), topic, fmt.Sprintf("Hello %d\n", i))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	success := false
	select {
	case <-ctx.Done():
		fmt.Println("context done")
		success = false
	case <-done:
		success = true
		fmt.Println("done")
	}

	fmt.Println("unsubscribing from broker")
	broker.Unsubscribe(topic, sub1)

	if success {
		fmt.Println("received all messages")
	} else {
		fmt.Println("failed to receive all messages")
	}
}

// Subscriber defines the public interface for receiving messages.
// The user provides a context to control blocking or cancellation.
type Subscriber interface {
	// Receive returns the message and a boolean indicating whether the message was successfully received.
	Receive(ctx context.Context) (string, bool)
}

type subscriber struct{}

func (s *subscriber) Receive(ctx context.Context) (string, bool) {
	return "", false
}

// Broker defines the interface for a publish-subscribe broker.
// It supports subscribing to topics, publishing messages, and unsubscribing.
type Broker interface {
	Subscribe(topic string) Subscriber
	Unsubscribe(topic string, sub Subscriber)
	Publish(ctx context.Context, topic, msg string)
}

type broker struct{}

func NewBroker() *broker {
	return &broker{}
}

func (b *broker) Subscribe(topic string) Subscriber {
	return &subscriber{}
}

func (b *broker) Unsubscribe(topic string, sub Subscriber) {}

func (b *broker) Publish(ctx context.Context, topic, msg string) {}
