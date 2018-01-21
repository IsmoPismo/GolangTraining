package main

import "fmt"

func main() {
	fmt.Println(greet("John", "Doe"))
}

func greet(name, surname string) (string, string) {
	return fmt.Sprint(name, " ", surname), fmt.Sprint(surname, " ", name)
}
