package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	minLatency = 10
	maxLatency = 5000
	timeout    = 3000
)

func main() {
	from := "Phnom Penh"
	to := "Bangkok"
	fmt.Println("Search operation started: ", from, to)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Millisecond)
	defer cancel()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("aborting due to interrupt...")
		cancel()
	}()

	res, err := Search(ctx, from, to)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	fmt.Println(res)
}

func Search(ctx context.Context, from, to string) ([]string, error) {
	res := make(chan []string)

	go func() {
		res <- slowSearch(from, to)
	}()

	// wait for 2 events: either one will be the result
	for {
		select {
		case dst := <-res:
			return dst, nil

		case <-ctx.Done():
			return []string{}, ctx.Err()
		}
	}
}

func slowSearch(from, to string) []string {
	britAir := fmt.Sprintf("%s-%s-british airways-11am", from, to)
	deltaAir := fmt.Sprintf("%s-%s-delta airways-12am", from, to)

	latency := rand.Intn(maxLatency-minLatency+1) - minLatency
	fmt.Println("latency added: ", latency)

	time.Sleep(time.Duration(latency) * time.Millisecond)

	return []string{britAir, deltaAir}
}
