package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/seikichi/gopl/ch02/ex02/lenconv"
	"github.com/seikichi/gopl/ch02/ex02/tempconv"
	"github.com/seikichi/gopl/ch02/ex02/weightconv"
)

func main() {
	if len(os.Args) == 1 {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			printValues(sc.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			printValues(arg)
		}
	}
}

func printValues(s string) {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	k := tempconv.Kelvin(t)
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s = %s\n%s = %s = %s\n%s = %s = %s\n",
		k, tempconv.KToC(k), tempconv.KToF(k),
		f, tempconv.FToK(f), tempconv.FToC(f),
		c, tempconv.CToK(c), tempconv.CToF(c))

	m := lenconv.Meter(t)
	ft := lenconv.Feet(t)
	fmt.Printf("%s = %s\n%s = %s\n", m, lenconv.MToF(m), ft, lenconv.FToM(ft))

	kg := weightconv.Kilogram(t)
	lb := weightconv.Pound(t)
	fmt.Printf("%s = %s\n%s = %s\n", kg, weightconv.KToP(kg), lb, weightconv.PToK(lb))

	fmt.Println()
}
