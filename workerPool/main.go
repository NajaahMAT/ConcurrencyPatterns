package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//Number of workers in the pool
	numWorkers := 3

	//Create a wait group to synchronize the completion of all workers
	var wg sync.WaitGroup

	//Create a channel for task distribution
	taskCh := make(chan Task)

	//Launch the workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(taskCh, &wg)
	}

	//Send the task to work pool
	tasks := generateTask()

	for _, task := range tasks {
		taskCh <- task
	}

	close(taskCh)

	// Wait for all workers to complete
	wg.Wait()

}

type Task struct {
	ID   int
	Data string
}

func generateTask() []Task {
	tasks := []Task{
		{ID: 1, Data: "Task 1"},
		{ID: 2, Data: "Task 2"},
		{ID: 3, Data: "Task 3"},
		{ID: 4, Data: "Task 4"},
		{ID: 5, Data: "Task 5"},
		{ID: 6, Data: "Task 6"},
		{ID: 7, Data: "Task 7"},
	}

	return tasks
}

func worker(taskCh <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskCh {
		time.Sleep(700 * time.Millisecond)
		result := fmt.Sprintf("Task ID: %d Data: %s", task.ID, task.Data)
		fmt.Println(result)
	}
}

//output
// Task ID: 3 Data: Task 3
// Task ID: 2 Data: Task 2
// Task ID: 1 Data: Task 1
// Task ID: 4 Data: Task 4
// Task ID: 5 Data: Task 5
// Task ID: 6 Data: Task 6
// Task ID: 7 Data: Task 7
