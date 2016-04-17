package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(join(",", os.Args...))
}

func join(sep string, a ...string) string {
	return strings.Join(a, sep)
}
