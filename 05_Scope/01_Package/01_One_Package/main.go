package main

import "fmt"

var x = 29

func main() {
	fmt.Print(x)
	foo()
}

func foo() {
	fmt.Print(x)
}
