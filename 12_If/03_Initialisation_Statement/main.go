package main

import "fmt"

func main() {
	b := true

	if food := "Pizza"; b {
		fmt.Print(food)
	}

	fmt.Print(food) // Won't run

}
