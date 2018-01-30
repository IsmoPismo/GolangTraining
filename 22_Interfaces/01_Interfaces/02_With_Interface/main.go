package main

import "fmt"

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{side: 10.55}
	fmt.Println("Area of Square with side", s.side, "is", s.area())
	info(s)
}
