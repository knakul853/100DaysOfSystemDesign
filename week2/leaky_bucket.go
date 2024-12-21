package main

import (
	"fmt"
	"time"
)

// LeakyBucket represents a rate limiter using the Leaky Bucket algorithm.
type LeakyBucket struct {
	capacity    int           // Maximum number of requests the bucket can hold.
	processRate time.Duration // Rate at which requests are processed.
	queue       chan struct{} // Queue to hold requests.
}

// NewLeakyBucket creates a new LeakyBucket with the given capacity and process rate.
func NewLeakyBucket(capacity int, processRate time.Duration) *LeakyBucket {
	bucket := &LeakyBucket{
		capacity:    capacity,
		processRate: processRate,
		queue:       make(chan struct{}, capacity),
	}

	// Start the leak process in a separate goroutine.
	go bucket.leak()

	return bucket
}

// leak processes requests from the bucket at the specified rate.
func (lb *LeakyBucket) leak() {
	for {
		select {
		case <-lb.queue:
			time.Sleep(lb.processRate) // Simulate processing time.
			fmt.Println("Request processed")
		}
	}
}

// Allow adds a request to the bucket if there is space.
func (lb *LeakyBucket) Allow() bool {
	select {
	case lb.queue <- struct{}{}:
		fmt.Println("Request added to bucket")
		return true
	default:
		fmt.Println("Bucket is full, request dropped")
		return false
	}
}

func Test() {
	// Create a leaky bucket with a capacity of 5 requests and a process rate of 1 request per second.
	bucket := NewLeakyBucket(5, time.Second)

	// Simulate 10 requests.
	for i := 1; i <= 10; i++ {
		bucket.Allow()
		time.Sleep(200 * time.Millisecond) // Simulate request interval.
	}

	// Wait to see the processing of requests.
	time.Sleep(10 * time.Second)
}
