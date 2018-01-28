package main

import "fmt"

type foo int

func main() {
	var x foo
	x = 42
	fmt.Printf("%T %v", x, x)
}
