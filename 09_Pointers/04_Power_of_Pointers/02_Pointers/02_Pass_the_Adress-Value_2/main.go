package main

import "fmt"

func zero(z *int) {
	*z = 0
	fmt.Println(z)
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
	fmt.Println(&x)
}
