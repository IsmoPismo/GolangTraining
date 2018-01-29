package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	first       string // None of these will be exported
	last        string
	age         int
	notExported int
}

func main() {
	ismar := person{First: "Ismar", Last: "Šaćirović", Age: 29, notExported: 26051988}
	fmt.Println(ismar)
	bs, _ := json.Marshal(ismar)
	fmt.Println(string(bs))
}
