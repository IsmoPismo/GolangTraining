package main

import "fmt"

func main() {
	var a int
	fmt.Println("Adress of a before: ", &a)
	fmt.Println("Value of a before: ", a)
	changeMe(&a)
	fmt.Println("Adress of a after: ", &a)
	fmt.Println("Value of a after: ", a)
}

func changeMe(z *int) {
	*z = 45
	fmt.Println("Adress of z: ", z)
	fmt.Println("Value of z: ", *z)
}
