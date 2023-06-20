package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	inputCh := make(chan int)
	outputCh := make(chan int)

	numWorkers := 3

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(inputCh, outputCh, &wg)
	}

	go func() {
		for _, val := range input {
			inputCh <- val
		}

		close(inputCh)
	}()

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	for val := range outputCh {
		fmt.Println(val)
	}

}

func worker(inputCh <-chan int, outputCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range inputCh {
		time.Sleep(500 * time.Millisecond)
		result := val * 2

		outputCh <- result
	}
}
