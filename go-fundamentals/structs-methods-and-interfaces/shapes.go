package structs_methods_and_interfaces

import "math"

// Shape interface
type Shape interface {
	// Computes the Perimeter of given Shape
	Perimeter() float64
	// Computes the Area of given Shape
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Computes the Perimeter of Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Computes the Area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Computes the Perimeter of Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Computes the Area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Represents any triangle
type Triangle struct {
	Sides [3]float64
}

// Computes the Perimeter of Triangle
func (t Triangle) Perimeter() float64 {
	sum := 0.0
	for _, side := range t.Sides {
		sum += side
	}
	return sum
}

// Computes the Area of Triangle using Heron's Formula
func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	acc := s

	for _, side := range t.Sides {
		acc *= (s - side)
	}
	return math.Sqrt(acc)
}
