package main

import "fmt"

func main() {
	var mapa = make(map[string]bool)
	mapa["Ismar"] = true
	mapa["Larisa"] = false

	for key, val := range mapa {
		fmt.Println(key, val)
	}
}
