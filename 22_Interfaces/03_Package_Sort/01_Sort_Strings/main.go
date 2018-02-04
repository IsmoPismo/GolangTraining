package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(s)
	// Witch basically turns s isto a 'StringSlice' and sorts it
	// { Sort(StringSlice(s)) }
	fmt.Println(s)
}
