package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	ismar := person{first: "Ismar", last: "Šaćirović"}
	fmt.Println(ismar.first)
}
