package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	ismar := person{First: "Ismar", Last: "Šaćirović", Age: 29, notExported: 26051988}
	bs, _ := json.Marshal(ismar)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
