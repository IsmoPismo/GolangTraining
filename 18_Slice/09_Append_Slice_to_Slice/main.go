package main

import "fmt"

func main() {
	one := []int{5, 5, 5}
	two := []int{6, 6, 6}
	one = append(one, two...)
	fmt.Println(one)
}
