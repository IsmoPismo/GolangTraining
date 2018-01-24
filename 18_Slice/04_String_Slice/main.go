package main

import "fmt"

func main() {

	stringSlice := []string{
		"hi!",
		"hey!",
		"yo!",
	}

	for i, currentEntry := range stringSlice {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(stringSlice); j++ {
		fmt.Println(j, stringSlice[j])
	}

}
