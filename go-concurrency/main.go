package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	started := time.Now()
	foods := []string{"mashroom pizza", "pasta", "kebab", "cake"}

	var wg sync.WaitGroup
	// or add waitgroup here
	// wg.Add(len(foods))

	for _, food := range foods {
		// add waitgroup per item in foods
		wg.Add(1)
		go func(f string) {
			cook(f)
			// complete waitgroup after cook function completed
			wg.Done()
		}(food)
	}

	// wait for all the concurrent function to finish with waitgroup
	wg.Wait()
	fmt.Printf("done in %v\n", time.Since(started))
}

func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", food)
	fmt.Println("")
}
