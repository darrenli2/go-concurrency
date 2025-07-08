package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Data string
}

func (t *Task) Execute() (error, string) {
	fmt.Println("Processing task", t.ID)
	time.Sleep(time.Second * 1)
	if t.ID%2 == 0 {
		return errors.New("task failed"), ""
	}
	return nil, t.Data
}

func processTasks(tasks []Task, concurrency int) {
	wg := sync.WaitGroup{}
	taskChan := make(chan Task, len(tasks))

	// Send all tasks to the channel
	for _, task := range tasks {
		taskChan <- task
	}
	close(taskChan)

	// Start worker goroutines
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskChan {
				err, result := task.Execute()
				if err != nil {
					fmt.Println("Error processing task", task.ID, err)
				} else {
					fmt.Println("Task", task.ID, "processed successfully with result", result)
				}
			}
		}()
	}

	wg.Wait()
}

func main() {
	tasks := []Task{
		{ID: 1, Data: "Task 1"},
		{ID: 2, Data: "Task 2"},
		{ID: 3, Data: "Task 3"},
		{ID: 4, Data: "Task 4"},
		{ID: 5, Data: "Task 5"},
	}
	processTasks(tasks, 2)
}
