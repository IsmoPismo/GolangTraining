package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := &Person{"Ismar", 29}
	fmt.Printf("Type %T \n", p1)   // *main.Person
	fmt.Printf("Value: %v \n", p1) // &{Ismar 29}
	fmt.Println(p1.Age)
	fmt.Println(p1.Name)
}
