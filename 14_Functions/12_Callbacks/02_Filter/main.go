package main

import "fmt"

func visit(sliceofint []int, callback func(int) bool) []int {
	var xs []int
	for _, number := range sliceofint {
		if callback(number) {
			xs = append(xs, number)
		}
	}
	return xs
}

func main() {
	funcExp := visit([]int{1, 2, 3, 5, 7}, func(n int) bool {
		return n > 2
	})
	fmt.Println(funcExp)
}
