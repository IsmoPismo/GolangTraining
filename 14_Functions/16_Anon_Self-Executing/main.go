package main

import "fmt"

func main() {
	func() {
		fmt.Println("I'm driving!")
	}() // Parens to execute
}
