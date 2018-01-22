package main

import "fmt"

func main() {
	c := fibonacci()
	fmt.Print(c)
}

func fibonacci() int {
	a := 1
	b := 2
	sum := 2
	var c int
	for sum < 4000000 {
		c = a + b
		a = b
		b = c
		if c%2 == 0 {
			sum += c
		}
	}
	return sum
}
