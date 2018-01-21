package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	fmt.Println(m) // Slice is a reference type
	chageSlice(m)  // That's why we the func changed
	fmt.Println(m) // the empty slice of a string
}

func chageSlice(xs []string) {
	xs[0] = "Ismar"
}
