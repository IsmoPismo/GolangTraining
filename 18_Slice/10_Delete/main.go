package main

import "fmt"

func main() {
	one := []string{"Element1", "Deleted Element", "Last Element"}
	one = append(one[:1], one[2:]...)
	fmt.Println(one)
}
