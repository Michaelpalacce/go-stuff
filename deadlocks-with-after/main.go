package main

import (
	"context"
	"fmt"
	"time"
)

func sleeper(ctx context.Context, result chan bool) {
	time.Sleep(time.Second * 2)
	select {
	case <-ctx.Done():
		println("Context is cancelled")
	default:
		fmt.Println("Returning!")
		result <- true
	}
}

func main() {
	resultChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go sleeper(ctx, resultChan)

	select {
	case result := <-resultChan:
		fmt.Printf("%v was returned!\n", result)
	case <-time.After(1 * time.Second):
		println("DEADLOCK AVOIDED!")
		cancel()
	}

	time.Sleep(time.Second * 5)
}
