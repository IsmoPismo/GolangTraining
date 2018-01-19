package main

import "fmt"

func main() {
	var largenum int
	var smallnum int
	fmt.Println("Enter large num: ")
	fmt.Scan(&largenum)
	fmt.Println("Enter small num: ")
	fmt.Scan(&smallnum)
	var remainder int = largenum % smallnum
	fmt.Printf("%v divided by %v gives the remainder of %v", largenum, smallnum, remainder)
}
