package main

import "fmt"

func main() {
	var word string
	changeString(&word)
	fmt.Println(word)
}

func changeString(w *string) {
	*w = "Chage the String"
}
