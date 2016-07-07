package main

import (
	"fmt"
	"time"
)

func ping(in <-chan struct{}, out chan<- struct{}, done <-chan struct{}, result chan<- int) {
	out <- struct{}{}
	count := 1

loop:
	for {
		select {
		case <-done:
			break loop
		default:
			out <- <-in
			count++
		}
	}
	result <- count
}

func pong(in <-chan struct{}, out chan<- struct{}) {
	for {
		out <- <-in
	}
}

func main() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	result := make(chan int)
	done := make(chan struct{})

	go func() {
		<-time.After(1 * time.Second)
		close(done)
	}()

	go pong(c2, c1)
	go ping(c1, c2, done, result)

	count := <-result
	fmt.Println(count)
}
