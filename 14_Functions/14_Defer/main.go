package main

import "fmt"

func main() {
	defer world() // Executes this before it exits the function
	hello()
}

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Print("world")
}
