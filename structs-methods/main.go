package main

import (
	"math"
)

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

func (rp Rectangle) Perimeter() float64 {
	perimeter := (rp.Width + rp.Height) * 2
	return perimeter
}
