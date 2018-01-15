package main

import "fmt"

var x = 0

func increment() int {
	x++
	return x
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(increment())
	}
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
