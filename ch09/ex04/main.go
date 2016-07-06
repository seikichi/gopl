package main

import "flag"

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

	in, out := pipeline(size)
	go func() { in <- struct{}{} }()
	<-out
}
