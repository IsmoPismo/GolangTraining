package main

import (
	"fmt"
)

func main() {
	var a [35]string
	fmt.Println(len(a))
	a[2] = "Number 23"
	fmt.Println(a[2])
	fmt.Printf("%T\t%T", a, a[2])
}
