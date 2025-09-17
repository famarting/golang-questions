package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := serveAPI(ctx, ":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

func serveAPI(ctx context.Context, address string) error {
	log.Printf("listening on %s", address)

	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	registerHandlers(mux)

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		log.Printf("shutting down")
		server.Shutdown(context.Background())
	}()

	return server.Serve(l)
}

// Register the handlers for the API.
func registerHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	})
}
