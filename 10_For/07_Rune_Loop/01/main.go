package main

import "fmt"

func main() {
	for i := 1; i <= 440; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
