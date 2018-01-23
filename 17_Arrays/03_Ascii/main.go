package main

import "fmt"

func main() {
	var x [61]string
	for i := 65; i <= 125; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
}
