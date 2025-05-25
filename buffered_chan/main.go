package main

import (
	"fmt"
	"time"
)

func worker(id int, tasks <-chan int, ready chan<- bool) {
	for task := range tasks {
		fmt.Printf("Worker %d started task %d\n", id, task)
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("Worker %d finished task %d\n", id, task)
		ready <- true
	}
}

func main() {
	tasks := make(chan int)
	ready := make(chan bool)

	go worker(1, tasks, ready)

	tasks <- 0

	for i := 1; i <= 3; i++ {
		<-ready
		fmt.Printf("Main: Worker is ready. Sending task %d\n", i)
		tasks <- i
	}

	<-ready
	close(tasks)
	fmt.Println("Main: All tasks sent and acknowledged.")
}
