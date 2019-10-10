package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var rate = time.Second

func gen() chan int {
	msgs := make(chan int)

	go func() {
		for {
			msgs <- rand.Intn(100000)
		}
	}()

	return msgs
}

func throttler(ctx context.Context, msgs chan int) chan int {
	out := make(chan int)

	go func() {
		for {
			select {
			case msg := <-msgs:
				out <- msg
				if rate > 0 {
					fmt.Printf("Limiting rate %v\n", rate)
					time.Sleep(rate)
				}
			case <-ctx.Done():
				fmt.Printf("Context done\n")
				return
			}
		}
	}()

	return out
}

func main() {

	ctx, c := context.WithCancel(context.Background())
	msgs := throttler(ctx, gen())

	go func() {
		for {
			<-msgs
		}
	}()

	rate = time.Millisecond
	time.Sleep(time.Second)
	rate = 100 * time.Millisecond
	time.Sleep(time.Second)
	rate = 1000 * time.Millisecond
	time.Sleep(time.Second)
	rate = 100 * time.Millisecond
	time.Sleep(time.Second)
	rate = 500 * time.Millisecond
	time.Sleep(time.Second)
	rate = 500 * time.Millisecond
	c()

	time.Sleep(time.Hour)

}
