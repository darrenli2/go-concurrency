package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Task represents a unit of work to be processed.
type Task struct {
	ID   int
	Data string
}

// Execute simulates processing the task, returning an error for even IDs.
func (t *Task) Execute() (error, string) {
	fmt.Printf("Processing task %d\n", t.ID)
	time.Sleep(time.Second) // Simulate work
	if t.ID%2 == 0 {
		return errors.New("task failed"), ""
	}
	return nil, t.Data
}

// processTasksWithSemaphore concurrently processes a slice of tasks with a given concurrency limit,
// using a buffered channel as a semaphore.
func processTasksWithSemaphore(tasks []Task, concurrency int) {
	var wg sync.WaitGroup
	// Create a buffered channel to act as a counting semaphore.
	// The capacity of the channel determines the maximum number of concurrent goroutines.
	semaphore := make(chan struct{}, concurrency)

	// Iterate over each task and launch a goroutine for it.
	for _, task := range tasks {
		// Acquire a "slot" from the semaphore.
		// This call will block if the semaphore channel is full (i.e., 'concurrency' goroutines
		// are already running), ensuring we don't exceed the limit.
		semaphore <- struct{}{} // Send an empty struct to occupy a slot

		wg.Add(1) // Increment the WaitGroup counter for each task goroutine
		// Use a closure to pass the task value to the goroutine
		go func(t Task) {
			defer wg.Done() // Decrement the counter when the goroutine exits

			// Release the "slot" back to the semaphore when the task is done.
			// This allows another blocked goroutine to proceed.
			defer func() { <-semaphore }() // Receive from the channel to free a slot

			// Execute the task and handle its result
			err, result := t.Execute()
			if err != nil {
				fmt.Printf("Error processing task %d: %v\n", t.ID, err)
			} else {
				fmt.Printf("Task %d processed successfully with result: %s\n", t.ID, result)
			}
		}(task) // Pass the current task to the goroutine
	}

	wg.Wait() // Wait for all task goroutines to complete their execution
	// It's good practice to close the semaphore channel if it's no longer needed,
	// though not strictly necessary here since the function exits.
	// close(semaphore)
}

func main() {
	// Define a slice of tasks to be processed.
	tasks := []Task{
		{ID: 1, Data: "Task 1 Data"},
		{ID: 2, Data: "Task 2 Data"},
		{ID: 3, Data: "Task 3 Data"},
		{ID: 4, Data: "Task 4 Data"},
		{ID: 5, Data: "Task 5 Data"},
		{ID: 6, Data: "Task 6 Data"},
		{ID: 7, Data: "Task 7 Data"},
		{ID: 8, Data: "Task 8 Data"},
		{ID: 9, Data: "Task 9 Data"},
		{ID: 10, Data: "Task 10 Data"},
	}
	fmt.Println("--- Processing tasks with a semaphore (concurrency limit of 3) ---")
	processTasksWithSemaphore(tasks, 3) // Process with a maximum of 3 concurrent workers
}
