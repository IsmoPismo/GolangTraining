package main

import "fmt"

// Square is a Shape with four equal straight sides and four right angles
type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type shape interface {
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{side: 10.55}
	info(s)
}
