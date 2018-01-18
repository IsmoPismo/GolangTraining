package main

import "fmt"

func main() {
	foo := 'Š'
	fmt.Printf("%v ima broj: ", string(foo))
	fmt.Print(foo)
	fmt.Printf(", type: %T. ", foo)
	fmt.Printf("U bajtovima: %v", []byte(string(foo)))
}
