package main

import "fmt"

type foo int

func main() {
	var x foo
	x = 42
	fmt.Printf("%T %v\n", x, x)

	var y int
	y = 21
	fmt.Printf("%T %v\n", y, y)

	// this doesn't work:
	// fmt.Println(x + y)

	// this conversion works:
	fmt.Println(int(x) - y - int(x))

}
