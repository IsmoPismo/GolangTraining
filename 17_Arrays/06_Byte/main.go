package main

import "fmt"

func main() {
	var x [30]byte

	for i := 0; i < 30; i++ {
		x[i] = byte(i)
	}

	for i, v := range x {
		fmt.Printf("value - %v\ttype - %T\tbinary - %b\n", v, v, v)
		if i == 29 {
			break
		}
	}
}
