package main

import "fmt"

func main() {
	// Example how hex takes up less chars than bin or dec
	for i := 1000000000; i < 1000000002; i++ {
		fmt.Printf("%d \t%b\t%#X\n", i, i, i)
	}

	// ASCII
	for i := 65; i < 121; i++ {
		fmt.Printf("%d \t %b \t %x \t %s\n", i, i, i, i)
	}
}
