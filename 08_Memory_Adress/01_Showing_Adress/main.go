package main

import "fmt"

func main() {
	a := 29

	fmt.Println("a - ", a)
	fmt.Println("a's memory Adress - ", &a)
	fmt.Printf("Adress in Decimal: %d\nAdress in Binary: %b", &a, &a)
}
