package main

import "fmt"

func main() {
	fmt.Println(square(10))
}

func square(n int) (ret int) {
	defer func() {
		recover()
		ret = n * n
	}()
	panic(nil)
}
