package main

import "fmt"

func main() {
	a := 29
	b := "Ismar"
	c := true
	d := 4.175

	fmt.Printf("%v \tType: %T \n", a, a)
	fmt.Printf("%v \tType: %T \n", b, b)
	fmt.Printf("%v \tType: %T \n", c, c)
	fmt.Printf("%v \tType: %T \n", d, d)

	fmt.Println("Zero-Init - Go Assigns Zero Value:")

	var e int
	var f string
	var g bool
	var h float64

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
}
