package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d: Online and waiting for jobs...\n", id)

	for job := range jobs {
		fmt.Printf("Worker %d: Received job %d\n", id, job)

		time.Sleep(time.Millisecond * 200)
		result := job * 10

		fmt.Printf("Worker %d: Processed job %d, sending result %d\n", id, job, result)
		results <- result
	}

	fmt.Printf("Worker %d: Jobs channel closed. All jobs processed. Shutting down.\n", id)
}

func main() {
	const numJobs = 9
	const bufferSize = 3

	jobs := make(chan int, bufferSize)
	results := make(chan int, bufferSize)

	var workerWg sync.WaitGroup

	workerWg.Add(2)
	go worker(1, jobs, results, &workerWg)
	go worker(2, jobs, results, &workerWg)

	var resultsCollectorWg sync.WaitGroup
	resultsCollectorWg.Add(1)
	go func() {
		defer resultsCollectorWg.Done()
		fmt.Println("ResultsCollector: Online and waiting for results...")
		for res := range results {
			fmt.Printf("ResultsCollector: Received result: %d\n", res)
		}
		fmt.Println("ResultsCollector: Results channel closed. All results received.")
	}()

	fmt.Println("Main: Sending jobs...")
	for j := 1; j <= numJobs; j++ {
		fmt.Printf("Main: Sending job %d to 'jobs' channel\n", j)
		jobs <- j
		if j%2 == 0 {
			time.Sleep(time.Millisecond * 50)
		}
	}
	fmt.Println("Main: All jobs have been sent to the 'jobs' channel.")

	close(jobs)
	fmt.Println("Main: 'jobs' channel closed.")

	fmt.Println("Main: Waiting for worker to complete all processing...")
	workerWg.Wait()
	fmt.Println("Main: Worker has completed.")

	close(results)
	fmt.Println("Main: 'results' channel closed.")

	fmt.Println("Main: Waiting for ResultsCollector to complete...")
	resultsCollectorWg.Wait()
	fmt.Println("Main: ResultsCollector has completed.")

	fmt.Println("Main: Program finished successfully.")
}
