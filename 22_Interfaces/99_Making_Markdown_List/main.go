package main

import (
	"fmt"
	"sort"
	"strings"
)

type markdown []string

func (md markdown) Len() int           { return len(md) }
func (md markdown) Swap(i, j int)      { md[i], md[j] = md[j], md[i] }
func (md markdown) Less(i, j int) bool { return md[i] < md[j] }

func main() {
	ss := []string{`01_First_Steps/     04_Variables/         07_Constants/      10_For/     13_Hands_on_Go/  16_Hands_on_Go_2/  19_Map/          22_Interfaces/
02_Numeral_System/  05_Scope/             08_Memory_Adress/  11_Switch/  14_Functions/    17_Arrays/         20_Hash_Tables/  99_Deck_of_Cards/
03_Package/         06_Blank_Identifier/  09_Pointers/       12_If/      15_Bool/         18_Slice/          21_Struct/
`}
	s := strings.Join([]string(ss), "")
	bs := []byte(s)
	end := strings.Split(string(bs), "/")
	md := markdown(end)
	for o, v := range md {
		md[o] = strings.Replace(strings.Replace(strings.TrimSpace(v)+"\n", "_", ". ", 1), "_", " ", -1)
	}
	sort.Strings(md)
	fmt.Println(md)
}
