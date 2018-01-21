package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting() // func expresion? Only when assigning to a variable?
	fmt.Printf("%T\n", greeting)
}
