package main

import "fmt"

func max(x int) int { // Declaring a func with a PARAMETER
	return 42 + x
}

func main() {
	max := max(7)    // Passing a ARGUMENT
	fmt.Println(max) // max is now the result, not the function
}

// don't do this; bad coding practice to shadow variables
