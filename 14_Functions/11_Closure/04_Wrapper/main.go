package main

import "fmt"

func wrapper() func() int {
	var x int
	fmt.Println(&x) // 0xc04200e080
	return func() int {
		x++
		return x
	}
}

func main() {
	x := wrapper()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Print(&x) // 0xc042004028
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
