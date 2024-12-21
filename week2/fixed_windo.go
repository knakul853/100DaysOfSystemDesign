package main

import (
	"fmt"
	"sync"
	"time"
)

// FixedWindow represents a rate limiter using the Fixed Window algorithm.
type FixedWindow struct {
	limit      int           // Maximum number of requests allowed in a window.
	windowSize time.Duration // Size of the window.
	count      int           // Current number of requests in the window.
	mutex      sync.Mutex    // Mutex to ensure thread-safe access.
	lastWindow time.Time     // Start time of the current window.
}

// NewFixedWindow creates a new FixedWindow with the given limit and window size.
func NewFixedWindow(limit int, windowSize time.Duration) *FixedWindow {
	return &FixedWindow{
		limit:      limit,
		windowSize: windowSize,
		lastWindow: time.Now(),
	}
}

// Allow checks if a request is allowed within the current window.
func (fw *FixedWindow) Allow() bool {
	fw.mutex.Lock()
	defer fw.mutex.Unlock()

	now := time.Now()
	if now.Sub(fw.lastWindow) >= fw.windowSize {
		// Move to a new window.
		fw.count = 0
		fw.lastWindow = now
	}

	if fw.count < fw.limit {
		fw.count++
		return true
	}
	return false
}

func TestFixedWindow() {
	// Create a fixed window with a limit of 5 requests per second.
	window := NewFixedWindow(5, time.Second)

	// Simulate 10 requests.
	for i := 1; i <= 10; i++ {
		if window.Allow() {
			fmt.Printf("Request %d: Allowed\n", i)
		} else {
			fmt.Printf("Request %d: Denied\n", i)
		}
		time.Sleep(200 * time.Millisecond) // Simulate request interval.
	}
}
