package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3)

	for i := 0; i < 13; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value:", mySlice[i])
	}
}
