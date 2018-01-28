package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type ououseven struct {
	Person
	First           string
	LicenseToMurder bool
}

func main() {
	p1 := ououseven{
		Person: Person{
			First: "Johny",
			Last:  "Knoxwill",
			Age:   23,
		},
		First:           "Jackass", // OVERWRT!
		LicenseToMurder: false,
	}

	fmt.Println(p1.First) // Aka "Jackass"
}
