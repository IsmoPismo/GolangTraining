package main

import "fmt"

// Square is a Shape with four equal straight sides and four right angles
type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func main() {
	s := Square{side: 10.55}
	fmt.Println("Area of Square with side", s.side, "is", s.area())
}
