package main

import (
	"fmt"
	"time"
)

func ping(in <-chan struct{}, out chan<- struct{}, result chan<- int, d time.Duration) {
	ticker := time.After(d)
	out <- struct{}{}
	count := 1

	for {
		select {
		case <-ticker:
			result <- count
			return
		default:
			out <- <-in
			count++
		}
	}
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

	go pong(c2, c1)
	go ping(c1, c2, result, 1*time.Second)

	count := <-result
	fmt.Println(count)
}
