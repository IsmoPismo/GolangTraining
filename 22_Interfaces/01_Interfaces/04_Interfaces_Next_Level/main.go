package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (z square) area() float64 {
	return z.side * z.side
}

func (z circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// A new method that takes the INTERFACE TYPE shape
func totalArea(ss ...shape) float64 {
	var allArea float64
	for _, s := range ss {
		allArea += s.area()
	}
	return allArea
}

func main() {
	s := square{side: 10.55}
	info(s)
	t := circle{radius: 0.73}
	info(t)
	fmt.Println(totalArea(s, t))

}
