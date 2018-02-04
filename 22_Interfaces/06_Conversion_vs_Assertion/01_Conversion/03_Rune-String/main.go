package main

import "fmt"

func main() {
	var a rune = 'r' // Rune aka Int32
	var b int32 = 'ž'
	fmt.Printf("%T %T\n", a, b)
	fmt.Println(a, b)
	fmt.Printf("%v %v", string(a), string(b))
}
