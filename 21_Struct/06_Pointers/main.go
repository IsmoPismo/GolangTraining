package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func main() {
	p1 := &person{"Ismar", 29}
	fmt.Printf("Type %T \n", p1)   // *main.person
	fmt.Printf("Value: %v \n", p1) // &{Ismar 29}
	fmt.Println(p1.Age)
	fmt.Println(p1.Name)
}
