package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	ismar := person{First: "Ismar", Last: "Šaćirović", Age: 29}
	bs, _ := json.Marshal(ismar)
	fmt.Println(string(bs))
}
