package main

import "fmt"

func main() {
	var a int
	changeMe(a)
	fmt.Println("Third ", a)
}

func changeMe(z int) {
	fmt.Println("First: ", z)
	z = 45
	fmt.Println("Second ", z)
}
