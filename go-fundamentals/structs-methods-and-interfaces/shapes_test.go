package structs_methods_and_interfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 36.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 62.831853072},
		{name: "Triangle", shape: Triangle{Sides: [3]float64{3, 4, 5}}, expected: 12.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if math.Abs(got-tt.expected) > 1e-5 {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.expected)
			}
		})

	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Sides: [3]float64{3, 4, 5}}, expected: 6.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if math.Abs(got-tt.expected) > 1e-5 {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.expected)
			}
		})

	}
}
