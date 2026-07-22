package main

import (
	"testing"
	"math"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, math.Pi * 10 * 10},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
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
