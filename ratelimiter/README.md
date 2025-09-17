# Ratelimiter

Design and implement a rate-limiting middleware for an HTTP API in Go. The API should only allow a certain number of requests from a single client within a specified time window. If the client exceeds the limit, subsequent requests should be rejected until the window resets.

Requirements:

Rate Limit: 10 requests per minute per client. The rate limit should be enforced by tracking client IP addresses. Once the rate limit is reached, return an HTTP 429 (Too Many Requests) response with a message indicating the wait time before the client can make more requests.