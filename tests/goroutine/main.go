package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks { // func worker bị block tại đây để chờ Task
		fmt.Printf("\nWorker %d xử lý task:: %d\n", id, task)
		// Xử lý task
	}
}

func main() {
	tasks := make(chan int, 100)
	var wg sync.WaitGroup

	// Khởi động 5 worker
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Gửi task
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)

	// Chờ tất cả worker hoàn thành
	wg.Wait()
}
