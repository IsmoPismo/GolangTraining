package main

import "fmt"

func greatest(list ...int) int {
	max := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			max = list[i]
		}
	}
	return max
}

func main() {
	fmt.Print(greatest(10, 8, 6, 4, 2, -5))
}
