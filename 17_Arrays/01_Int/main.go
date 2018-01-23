package main

import (
	"fmt"
)

func main() {
	var a [35]int
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(a[2])
	a[2] = 125
	fmt.Println(a[2])
}
