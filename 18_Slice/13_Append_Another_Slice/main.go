package main

import "fmt"

func main() {
	oneSlice := []int{2, 3, 4}
	another := []int{5, 8}

	oneSlice = append(another, 8)
	fmt.Print(oneSlice)
}
