package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type DoubleOhSeven struct {
	Person
	LicenseToMurder bool
}

func (p Person) Greeting() {
	fmt.Println("I'm a regular guy...")
}

func (a DoubleOhSeven) Greeting() {
	fmt.Println("Bond...\tJames Bond")
}

func main() {
	p1 := Person{
		Name: "Johny",
	}

	p2 := DoubleOhSeven{
		Person: Person{
			Name: "James Bond",
		},
		LicenseToMurder: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
