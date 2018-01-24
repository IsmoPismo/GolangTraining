package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 8}

	for index, value := range mySlice {
		fmt.Println(index, value)
	}
}
