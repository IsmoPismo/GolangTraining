package main

import "fmt"

func main() {
	switch "Todd" {
	case "Casey", "Todd":
		fmt.Println("Wasap Casey or, err Todd")
	case "Pinđa", "Dennis":
		fmt.Println("Hi Pinđa or.. Dennis?")
	default:
		fmt.Println("You must be new")
	}
}
