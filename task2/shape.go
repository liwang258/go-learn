package task2

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length uint
	Width  uint
}

type Circle struct {
	Radius uint
}

func (r *Rectangle) Area() float64 {
	return float64(r.Length * r.Width)
}

func (c *Circle) Area() float64 {
	return math.Pi * float64(c.Radius*c.Radius)
}

func (r *Rectangle) Perimeter() float64 {
	return float64(2 * (r.Length + r.Width))
}

func (c *Circle) Perimeter() float64 {
	return float64(2 * math.Pi * float64(c.Radius))
}
