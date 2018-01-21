package main

import "fmt"

func main() {
	fmt.Println(avarage(42, 58, 95, 105, 23, 88))
}

func avarage(params ...float64) float64 {
	fmt.Printf("%v - %T \n", params, params) // [42 58 95 105 23 88] => a slice
	// []float => type: slice of float64
	var sum float64
	for _, v := range params {
		sum += v
	}
	return sum / float64(len(params))
}
