package main

import (
	"flag"
	"fmt"
	"os"
)

func pipeline(size int) (chan<- struct{}, <-chan struct{}) {
	var in chan<- struct{}
	var out <-chan struct{}

	for i := 0; i < size; i++ {
		c := make(chan struct{})
		go func(in chan<- struct{}, out <-chan struct{}) {
			for {
				in <- <-out
			}
		}(c, out)

		if in == nil {
			in = c
		}
		out = c
	}
	return in, out
}

func main() {
	size := 100
	flag.IntVar(&size, "size", 8000, "pipeline size")
	flag.Parse()

	fmt.Fprintf(os.Stderr, "Create a pipeline (size = %d) ...\n", size)
	in, out := pipeline(size)

	fmt.Fprintf(os.Stderr, "Send `struct{}{}` to the pipeline (size = %d) ...\n", size)
	go func() { in <- struct{}{} }()
	<-out
}
