package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First string
	Bday  int
}

func main() {
	p1 := person{"Ike' Casiljas", 1051989}
	json.NewEncoder(os.Stdout).Encode(p1)
}
