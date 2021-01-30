package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}
type rectangle struct {
	w float64
	h float64
}
type circle struct {
	r float64
}

func (r rectangle) Area() float64 {
	return r.w * r.h
}

func (r rectangle) Perimeter() float64 {
	return 2*r.w + 2*r.h
}

func (c circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func main() {
	var shapes []shape
	shapes = make([]shape, 0)
	shapes = append(shapes, rectangle{3, 4})
	shapes = append(shapes, circle{4})
	for _, s := range shapes {
		fmt.Printf("Area is %v\n", s.Area())
		fmt.Printf("Perimeter is %v\n", s.Perimeter())
	}
}
