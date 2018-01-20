package main

import "fmt"

func main() {
	greet("Jane", "Doe")
	greet("John", "Doe")
}

func greet(name, surname string) {
	fmt.Println(name, surname)
}
