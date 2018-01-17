package main

import "fmt"

func main() {
	a := "Diggidy"
	fmt.Println(a)
	fmt.Println(&a)

	var b *string = &a
	fmt.Println(b)
	// Dereferencing
	fmt.Println(*b)

}
