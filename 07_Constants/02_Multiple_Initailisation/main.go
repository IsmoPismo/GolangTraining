package main

import "fmt"

const (
	Pi       = 3.14
	language = "Go"
)

func main() {
	const q = 101

	fmt.Println("pi is ", Pi)
	fmt.Printf("This lang is " + language)
}
