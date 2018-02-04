package main

import (
	"fmt"
)

type person struct {
	Name string
}

type doubleOhSeven struct {
	person
	LicenseToMurder bool
}

func (p person) Greeting() {
	fmt.Println("I'm a regular guy...")
}

func (a doubleOhSeven) Greeting() {
	fmt.Println("Bond...\tJames Bond")
}

func main() {
	p1 := person{
		Name: "Johny",
	}

	p2 := doubleOhSeven{
		person: person{
			Name: "James Bond",
		},
		LicenseToMurder: true,
	}

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
