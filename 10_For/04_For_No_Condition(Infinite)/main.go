package main

import "fmt"

func main() {
	i := 1
	for {
		fmt.Print(i, ", ") // CTRL + C to stop
		i++
	}
}
