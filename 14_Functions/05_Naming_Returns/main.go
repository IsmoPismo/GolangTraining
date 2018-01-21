package main

import "fmt"

func main() {
	fmt.Println(greet("John", "Doe"))
}

func greet(name, surname string) (s string) {
	s = fmt.Sprint(name, surname)
	return
}
