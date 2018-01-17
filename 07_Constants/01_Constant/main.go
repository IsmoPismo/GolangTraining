package main

import "fmt"

const p string = "Prasa je zdrava"

// const x int (not posible)

func main() {
	const q = 101

	fmt.Println("p - ", p)
	fmt.Printf("q - %v", q)
}
