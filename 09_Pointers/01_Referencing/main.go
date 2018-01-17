package main

import "fmt"

func main() {
	a := "Diggidy"
	fmt.Println(a)
	fmt.Println(&a)

	var b *string = &a
	fmt.Println(b)
}

// b is type pointer -> Points to a Memory Adress where "Diggidy" is stored
// *string - star is part of the type
