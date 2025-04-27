package main

import "fmt"

// Refer to: https://refactoring.guru/design-patterns/bridge
/*
Bridge is a structural design pattern that
lets you split a large class or a set of closely related classes into
two separate hierarchies—abstraction and implementation—
which can be developed independently of each other.
*/

// Renderer is the Implementor interface
type Renderer interface {
	RenderCircle(radius float64)
}

// VectorRenderer is a Concrete Implementor
type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing a circle with radius %.2f using Vector Renderer\n", radius)
}

// RasterRenderer is another Concrete Implementor
type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius float64) {
	fmt.Printf("Drawing a circle with radius %.2f using Raster Renderer\n", radius)
}

// Shape is the Abstraction
type Shape interface {
	Draw()
}

// Circle is a Refined Abstraction
type Circle struct {
	Radius   float64
	Renderer Renderer
}

func (c *Circle) Draw() {
	c.Renderer.RenderCircle(c.Radius)
}

func main() {
	// Use VectorRenderer with Circle
	vectorRenderer := &VectorRenderer{}
	circle1 := &Circle{Radius: 5, Renderer: vectorRenderer}
	circle1.Draw()

	// Use RasterRenderer with Circle
	rasterRenderer := &RasterRenderer{}
	circle2 := &Circle{Radius: 10, Renderer: rasterRenderer}
	circle2.Draw()
}
