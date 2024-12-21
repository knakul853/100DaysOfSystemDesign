package main

import (
	"fmt"
	"sync"
	"time"
)

// TokenBucket represents a rate limiter using the Token Bucket algorithm.
type TokenBucket struct {
	capacity   int           // Maximum number of tokens the bucket can hold.
	tokens     int           // Current number of tokens in the bucket.
	refillRate time.Duration // Rate at which tokens are added to the bucket.
	lastRefill time.Time     // Last time the bucket was refilled.
	mutex      sync.Mutex    // Mutex to ensure thread-safe access.
}

// NewTokenBucket creates a new TokenBucket with the given capacity and refill rate.
func NewTokenBucket(capacity int, refillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity, // Start with a full bucket.
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// Refill adds tokens to the bucket based on the elapsed time since the last refill.
func (tb *TokenBucket) Refill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastRefill)

	// Calculate the number of tokens to add.
	tokensToAdd := int(elapsed / tb.refillRate)
	if tokensToAdd > 0 {
		tb.tokens = tb.tokens + tokensToAdd
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
		tb.lastRefill = now
	}
}

// Allow checks if a token is available and consumes it if possible.
func (tb *TokenBucket) Allow() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	tb.Refill() // Refill the bucket before checking.
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func main() {
	// Create a token bucket with a capacity of 10 tokens and a refill rate of 1 token per second.
	bucket := NewTokenBucket(10, time.Second)

	// Simulate 20 requests.
	for i := 1; i <= 20; i++ {
		if bucket.Allow() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Denied\n", i)
		}
		time.Sleep(100 * time.Millisecond) // Simulate request interval.
	}
}
