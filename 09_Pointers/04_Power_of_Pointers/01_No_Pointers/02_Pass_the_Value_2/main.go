package main

import "fmt"

func zero(z int) {
	z = 0
	fmt.Printf("%p", &z) // Another way of showing adress (p for pointer?)
}

func main() {
	x := 5
	zero(x)
	fmt.Printf("%p", &x)
}
