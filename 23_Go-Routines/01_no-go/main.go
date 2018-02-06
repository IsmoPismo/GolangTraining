package main

import "fmt"

func main() {
	foo()
	bar()
}

func foo() {
	for i := 0; i < 223; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 3; i++ {
		fmt.Println("Bar:", i)
	}
}
