package main

import (
	"testing"
	"math"
)

// func TestArea(t *testing.T) {

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		got := rectangle.Area()
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := math.Pi * 10 * 10

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }

func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: math.Pi * 10 * 10},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the 't.Run' test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape Perimeter
		want float64
	}{
		{Rectangle{14, 7}, 42.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
