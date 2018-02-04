package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter your number: ")
	fmt.Scan(&num)
	fmt.Print(half(num))
}

func half(num int) (int, bool) {
	if num%2 == 0 {
		return num / 2, true
	}
	return num / 2, false
}
