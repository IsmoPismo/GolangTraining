package main

import "fmt"

func main() {
	// Set up the pipeline.
	c := gen(2, 4)
	out := sq(c)

	// Consume the output.
	fmt.Println(<-out) // 4
	fmt.Println(<-out) // 9
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(input chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			out <- n * n
		}
		close(out)
	}()
	return out
}
