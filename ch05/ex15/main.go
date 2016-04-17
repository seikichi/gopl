package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	vals := []int{}
	for _, v := range os.Args[1:] {
		if i, err := strconv.Atoi(v); err == nil {
			vals = append(vals, i)
		}
	}
	m, err := max(vals...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("max(%v) = %v\n", vals, m)
}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("invalid arguments")
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if m < v {
			m = v
		}
	}
	return m, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("invalid arguments")
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if m > v {
			m = v
		}
	}
	return m, nil
}

func max2(val int, vals ...int) int {
	m := val
	for _, v := range vals {
		if m < v {
			m = v
		}
	}
	return m
}

func min2(val int, vals ...int) int {
	m := val
	for _, v := range vals {
		if m > v {
			m = v
		}
	}
	return m
}
