package main

import (
	"math"
)

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (c Circle) Area() float64 {
	area := c.Radius * c.Radius * math.Pi
	return area
}

func (r Rectangle) Area() float64 {
	area := r.Width * r.Height
	return area
}

func (t Triangle) Area() float64 {
	area := (t.Base * t.Height) / 2
	return area
}
