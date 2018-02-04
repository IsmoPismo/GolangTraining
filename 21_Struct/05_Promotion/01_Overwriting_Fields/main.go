package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type ououseven struct {
	person
	First           string
	LicenseToMurder bool
}

func main() {
	p1 := ououseven{
		person: person{
			First: "Johny",
			Last:  "Knoxwill",
			Age:   23,
		},
		First:           "Jackass", // OVERWRT!
		LicenseToMurder: false,
	}

	fmt.Println(p1.First) // Aka "Jackass"
}
