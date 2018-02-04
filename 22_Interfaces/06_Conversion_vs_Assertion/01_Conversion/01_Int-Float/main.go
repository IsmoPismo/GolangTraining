package main

import "fmt"

func main() {
	x := 14
	y := 14.1
	fmt.Printf("%T %T\n", x, y)
	fmt.Println(float64(x) + y)
}
