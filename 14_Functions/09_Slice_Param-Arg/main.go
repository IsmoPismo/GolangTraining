package main

import "fmt"

func main() {
	data := []float64{42, 51, 95, 105, 23, 88}
	fmt.Println(avarage(data))
}

func avarage(params []float64) float64 {
	var sum float64
	for _, v := range params {
		sum += v
	}
	return sum / float64(len(params))
}
