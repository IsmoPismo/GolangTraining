package main

import "fmt"

func main() {
	data := []float64{42, 58, 95, 105, 23, 88}
	fmt.Printf("data is type %T \n", data)
	n := avarage(data...)
	fmt.Printf("Avarage is %g \n", n)
}

func avarage(params ...float64) float64 {
	var sum float64
	for _, v := range params {
		sum += v
	}
	return sum / float64(len(params))
}
