package main

import "fmt"

// Not in Block
func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}

var y = 42
