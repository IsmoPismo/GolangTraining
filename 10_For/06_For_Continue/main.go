package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 1 {
			continue // dont print odd number
		}
		fmt.Println(i)
		if i >= 27 {
			break
		}
	}
}
