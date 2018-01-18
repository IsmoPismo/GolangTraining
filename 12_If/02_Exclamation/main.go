package main

import "fmt"

func main() {
	if !true {
		fmt.Print("This did not run")
	}

	if !false {
		fmt.Print("This ran")
	}
}
