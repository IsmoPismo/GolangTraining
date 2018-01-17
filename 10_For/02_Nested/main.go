package main

import "fmt"

func main() {
	for i := 1; i < 7; i += 2 {
		fmt.Println(i)
		for j := 2; j < 9; j += 2 {
			fmt.Println("Inside: ", j)
		}
	}

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}
