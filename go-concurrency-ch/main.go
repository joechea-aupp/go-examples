package main

import (
	"fmt"
	"time"
)

func main() {
	started := time.Now()
	foods := []string{"mashroom pizza", "pasta", "kebab", "cake"}

	// buffer channel, by default it has the capacity of 1, meaning it can only be send 1 item to a channel at a time
	results := make(chan bool)
	// results <-true where true is a value to send to a channel
	// to recieve a value, we can pull data from the variable
	// var := <-result or just <-result if we don't want to use the result as value later
	// if we don't pick and new item in a channel, it will be full and can't be populate in any other item until it's available again.

	for _, food := range foods {
		go func(f string) {
			cook(f)
			results <- true
		}(food)
	}

	for i := 0; i < len(foods); i++ {
		// access empty channel value will be block, until the value is provided and move to another cycle
		<-results
	}

	fmt.Printf("done in %v\n", time.Since(started))
}

func cook(food string) {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	fmt.Printf("done cooking %s\n", food)
	fmt.Println("")
}
