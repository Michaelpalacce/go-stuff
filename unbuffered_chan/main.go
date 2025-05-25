package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func eventProducer(id int, events chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 3 {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		event := fmt.Sprintf("Event from Producer %d: #%d", id, i)
		fmt.Printf("Producer %d sending: %s\n", id, event)
		events <- event
	}
}

func eventConsumer(events <-chan string, done chan<- bool) {
	for event := range events {
		fmt.Printf("Consumer processing: %s\n", event)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("Consumer finished.")
	done <- true
}

func main() {
	eventChannel := make(chan string, 5)

	var wg sync.WaitGroup
	numProducers := 3

	consumerDone := make(chan bool)
	go eventConsumer(eventChannel, consumerDone)

	for i := range numProducers {
		wg.Add(1)
		go eventProducer(i, eventChannel, &wg)
	}

	wg.Wait()
	fmt.Println("All producers finished sending.")

	close(eventChannel)

	<-consumerDone
	fmt.Println("Application finished.")
}
