package main

import "fmt"

// Reference: https://refactoring.guru/design-patterns/prototype
/* Prototype is a creational design pattern that
lets you copy existing objects without making your code dependent on their classes.
*/

// Shape interface
type Shape interface {
	Clone() Shape
	GetInfo() string
}

// Circle struct
type Circle struct {
	Radius int
	Color  string
}

// Clone method for Circle
func (c *Circle) Clone() Shape {
	return &Circle{
		Radius: c.Radius,
		Color:  c.Color,
	}
}

// GetInfo returns Circle details
func (c *Circle) GetInfo() string {
	return fmt.Sprintf("Circle: Radius=%d, Color=%s", c.Radius, c.Color)
}

// Rectangle struct
type Rectangle struct {
	Width  int
	Height int
	Color  string
}

// Clone method for Rectangle
func (r *Rectangle) Clone() Shape {
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
		Color:  r.Color,
	}
}

// GetInfo returns Rectangle details
func (r *Rectangle) GetInfo() string {
	return fmt.Sprintf("Rectangle: Width=%d, Height=%d, Color=%s", r.Width, r.Height, r.Color)
}

func main() {
	// Create original shapes
	circle1 := &Circle{Radius: 10, Color: "Red"}
	rectangle1 := &Rectangle{Width: 20, Height: 10, Color: "Blue"}

	// Clone shapes
	circle2 := circle1.Clone()
	rectangle2 := rectangle1.Clone()

	// Print original and cloned objects
	fmt.Println("Original:", circle1.GetInfo())
	fmt.Println("Cloned  :", circle2.GetInfo())

	fmt.Println("Original:", rectangle1.GetInfo())
	fmt.Println("Cloned  :", rectangle2.GetInfo())
}
