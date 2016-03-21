package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	rotate(a, 2)
	fmt.Println(a)
}

func rotate(s []int, d int) {
	// return if s is empty
	n := len(s)
	if n == 0 {
		return
	}

	// normalize move distance
	d = d % n
	if d < 0 {
		d += n
	}
	// compute GCD(d, n) for jaggling algorithm
	g := gcd(d, n)

	// start jaggling algorithm
	for i := 0; i < g; i++ {
		si := s[i]
		j := i
		for {
			k := j + d
			if k >= n {
				k = k - n
			}
			if k == i {
				break
			}
			s[j] = s[k]
			j = k
		}
		s[j] = si
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
