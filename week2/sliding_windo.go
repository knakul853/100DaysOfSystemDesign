package main

import (
	"fmt"
	"sync"
	"time"
)

// SlidingWindow represents a rate limiter using the Sliding Window algorithm.
type SlidingWindow struct {
	limit      int           // Maximum number of requests allowed in the sliding window.
	windowSize time.Duration // Size of the sliding window.
	requests   []time.Time   // List of request timestamps in the current window.
	mutex      sync.Mutex    // Mutex to ensure thread-safe access.
}

// NewSlidingWindow creates a new SlidingWindow with the given limit and window size.
func NewSlidingWindow(limit int, windowSize time.Duration) *SlidingWindow {
	return &SlidingWindow{
		limit:      limit,
		windowSize: windowSize,
		requests:   make([]time.Time, 0),
	}
}

// Allow checks if a request is allowed within the sliding window.
func (sw *SlidingWindow) Allow() bool {
	sw.mutex.Lock()
	defer sw.mutex.Unlock()

	now := time.Now()
	// Remove expired requests.
	expired := now.Add(-sw.windowSize)
	newRequests := make([]time.Time, 0)
	for _, req := range sw.requests {
		if req.After(expired) {
			newRequests = append(newRequests, req)
		}
	}
	sw.requests = newRequests

	// Check if the request limit is exceeded.
	if len(sw.requests) < sw.limit {
		sw.requests = append(sw.requests, now)
		return true
	}
	return false
}

func TestSlidingWindow() {
	// Create a sliding window with a limit of 5 requests in the last 10 seconds.
	window := NewSlidingWindow(5, 10*time.Second)

	// Simulate 10 requests.
	for i := 1; i <= 10; i++ {
		if window.Allow() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Denied\n", i)
		}
		time.Sleep(2 * time.Second) // Simulate request interval.
	}
}
