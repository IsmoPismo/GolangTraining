package main

import "fmt"

func main() {
	var x int
	incr := func() int {
		x++
		return x
	}

	fmt.Println(incr())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
