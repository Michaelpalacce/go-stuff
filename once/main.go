package main

import (
	"fmt"
	"sync"
)

func main() {
	init := sync.Once{}

	for range 5 {
		init.Do(func() {
			fmt.Println("Hey")
		})
	}
}

