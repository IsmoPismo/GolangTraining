package main

import "fmt"

const (
	_ = iota      // 0
	A = iota * 10 // 1 * 10
	B = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
}
