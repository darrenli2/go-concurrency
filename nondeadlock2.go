package main

import (
	"fmt"
	"sync" // Import the sync package for WaitGroup
	"time" // Added for demonstration of concurrent execution
)

func main() {
	var wg sync.WaitGroup // Declare a WaitGroup

	for msg := range []int{1, 2, 3} {
		done := make(chan bool) // Unbuffered channel for synchronization within each pair

		// Increment the WaitGroup counter for each goroutine we launch
		wg.Add(2) // Adding 2 because we launch 'run' and 'executeMethod' goroutines

		go func(m int, d chan bool) {
			defer wg.Done() // Decrement counter when 'run' goroutine finishes
			run(m, d)
		}(msg, done)

		go func(m int, d chan bool) {
			defer wg.Done() // Decrement counter when 'executeMethod' goroutine finishes
			executeMethod(m, d)
		}(msg, done)

		// The 'done' channel is used for synchronization *between* run and executeMethod
		// It's not directly used by main to wait for them here, but rather
		// the WaitGroup handles the overall program waiting.
		// The original commented out `<-done` would only wait for executeMethod,
		// and not the run goroutine.
	}

	// Wait for all goroutines launched in the loop to complete
	fmt.Println("Main goroutine waiting for all other goroutines to finish...")
	wg.Wait()
	fmt.Println("All goroutines finished. Main goroutine exiting.")
}

func run(msg int, done chan bool) {
	// This goroutine will block until executeMethod sends a value
	val := <-done
	fmt.Printf("run (msg: %d) received: %t\n", msg, val)
}

func executeMethod(r int, done chan bool) {
	arr := []string{"a", "b", "c"}
	for _, val := range arr {
		fmt.Printf("iteration is %d (msg: %d), value is %s\n", r, r, val)
		time.Sleep(100 * time.Millisecond) // Added a small delay to make execution order more apparent
	}
	done <- true // Send a signal to the 'run' goroutine
}
