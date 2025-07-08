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

// processTasks concurrently processes a slice of tasks with a given concurrency limit.
func processTasks(tasks []Task, concurrency int) {
	var wg sync.WaitGroup
	// Create an unbuffered channel for tasks. An unbuffered channel is suitable
	// here because all tasks are sent to the channel before workers start,
	// and the channel is closed immediately after all tasks are sent.
	taskChan := make(chan Task)

	// Start worker goroutines. These goroutines will read tasks from taskChan
	// until the channel is closed.
	for i := 0; i < concurrency; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each worker
		go func() {
			defer wg.Done() // Decrement the counter when the goroutine exits
			for task := range taskChan {
				// Execute the task and handle its result
				err, result := task.Execute()
				if err != nil {
					fmt.Printf("Error processing task %d: %v\n", task.ID, err)
				} else {
					fmt.Printf("Task %d processed successfully with result: %s\n", task.ID, result)
				}
			}
		}()
	}

	// Send all tasks to the channel. This must happen after workers are started
	// to avoid deadlocks if the channel were unbuffered and no one was reading.
	for _, task := range tasks {
		taskChan <- task
	}
	// Close the channel after all tasks have been sent. This signals to the
	// worker goroutines that no more tasks will arrive, allowing them to exit
	// their `for range` loops gracefully.
	close(taskChan)

	wg.Wait() // Wait for all worker goroutines to complete their execution
}

func main() {
	// Define a slice of tasks to be processed.
	tasks := []Task{
		{ID: 1, Data: "Task 1 Data"},
		{ID: 2, Data: "Task 2 Data"},
		{ID: 3, Data: "Task 3 Data"},
		{ID: 4, Data: "Task 4 Data"},
		{ID: 5, Data: "Task 5 Data"},
	}
	// Process the tasks with a maximum of 2 concurrent workers.
	processTasks(tasks, 2)
}
