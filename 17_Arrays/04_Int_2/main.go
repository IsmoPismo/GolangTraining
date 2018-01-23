package main

import "fmt"

func main() {
	var x [61]int
	fmt.Println(x)
	fmt.Println(len(x))
	x[22] = 777
	fmt.Println(x)
	fmt.Println(len(x))
}
