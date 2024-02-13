package main

import (
	"fmt"
)

// Shape is an interface for geometric shapes

// Rectangle represents a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter calculates the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Square represents a square
type Square struct {
	SideLength float64
}

// Area calculates the area of the square
func (s Square) Area() float64 {
	return s.SideLength * s.SideLength
}

// Perimeter calculates the perimeter of the square
func (s Square) Perimeter() float64 {
	return 4 * s.SideLength
}

func main() {
	// Create a rectangle and a square
	rectangle := Rectangle{Width: 3, Height: 4}
	square := Square{SideLength: 5}

	// Calculate and print area and perimeter of each shape
	printShapeDetails(rectangle)
	printShapeDetails(square)
}

// printShapeDetails prints the area and perimeter of a shape
func printShapeDetails(shape Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
}
