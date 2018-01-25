package main

import "fmt"

func main() {
	myGreetings := map[int]string{
		1: "Đesi ba",
		2: "Šta ima?",
		3: "Kako si?",
	}

	if value, exists := myGreetings[2]; exists {
		fmt.Println("A Value Exists in that position")
		fmt.Println(`Tha Value is:`, value)
		fmt.Println("exists:", exists)
	}
}
