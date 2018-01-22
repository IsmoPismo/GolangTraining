package main

import "fmt"

func main() {
	var a int
	changeMe(&a)
}

func changeMe(z *int) {
	fmt.Printf("Adress of z: %d \n", z)
	y := z // In conclusion, a assignment assigns a different address (as expected)
	fmt.Printf("%d", &y)
}
