package main

import "fmt"

func main() {
	h, even := half(5)
	fmt.Println(h, even) // Println prints all items in the same line
}

func half(num int) (int, bool) {
	return num / 2, num%2 == 0
}
