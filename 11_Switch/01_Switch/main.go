package main

import "fmt"

func main() {
	switch "Todd" {
	case "Casey":
		fmt.Println("Wasap Casey")
	case "Tod":
		fmt.Println("Hi Todd")
	default:
		fmt.Println("You must be new")
		var name string
		fmt.Scan(&name)
		fmt.Println("Hello ", name)
	}
}
