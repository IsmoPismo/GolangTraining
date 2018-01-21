package main

import "fmt"

func visit(sliceofint []int, callback func(int)) {
	for _, number := range sliceofint {
		callback(number)
	}
}

func main() {
	visit([]int{1, 2, 3, 5, 7}, func(n int) {
		fmt.Print(n)
	})
}
