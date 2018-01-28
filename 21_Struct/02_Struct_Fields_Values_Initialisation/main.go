package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 75}
	p2 := person{"Gareth", "Bale", 29}
	fmt.Println(p1, p2)

	// Print Field-Values
	fmt.Println(p1.age, p1.first, p1.last)
}
