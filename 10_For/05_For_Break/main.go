package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Print(i)
		if i >= 10 {
			fmt.Print(" ... The End")
			break
		} else {
			fmt.Print(", ")
		}
		i++
	}
}
