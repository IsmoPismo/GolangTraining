package main

import "fmt"

func main() {
	h, even := half(5)
	fmt.Println(h, even)
}

func half(num int) (float64, bool) {
	return float64(num) / 2, num%2 == 0 // Convert int to float64 for return
}
