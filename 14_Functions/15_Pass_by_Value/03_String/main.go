package main

import "fmt"

func main() {
	var a string
	changeMe(a)
	fmt.Println("Third: ", a)
}

func changeMe(z string) {
	fmt.Println("First: ", z)
	z = "45"
	fmt.Println("Second ", z)
}
