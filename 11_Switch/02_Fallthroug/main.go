package main

import "fmt"

func main() {
	switch "Todd" {
	case "Casey":
		fmt.Println("Wasap Casey")
	case "Todd":
		fmt.Println("Hi Todd")
		fallthrough // Ide dole
	case "Pinđa":
		fmt.Println("Hi Pinđa")
		fallthrough // Ide dole
	case "Denis":
		fmt.Println("Hi Dennis")
	default:
		fmt.Println("You must be new")
	}
}
