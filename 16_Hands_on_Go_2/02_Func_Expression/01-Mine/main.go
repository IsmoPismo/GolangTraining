package main

import "fmt"

func main() {
	// User Input
	var num int
	fmt.Print("Enter your number: ")
	fmt.Scan(&num)
	//Func Expresion
	half := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}
	// Output
	fmt.Print(half(num))
}
