package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var pile person
	fmt.Println(pile.Age)

	bs := []byte(`{"First":"Larisa","Last":"Pile","Age":23}`)
	json.Unmarshal(bs, &pile)

	fmt.Println(pile.Age)
}
