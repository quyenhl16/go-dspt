// Reference: https://refactoring.guru/design-patterns/adapter
// Adapter is a structural design pattern that allows objects with incompatible interfaces to collaborate.

package main

import (
	"fmt"
	"math"
)

// RoundHole expects RoundPeg
type RoundHole struct {
	Radius float64
}

// Fits checks if the peg fits into the hole
func (r *RoundHole) Fits(peg RoundPeg) bool {
	return peg.GetRadius() <= r.Radius
}

// RoundPeg interface (expected by RoundHole)
type RoundPeg interface {
	GetRadius() float64
}

// SimpleRoundPeg Simple RoundPeg struct
type SimpleRoundPeg struct {
	Radius float64
}

// GetRadius method for SimpleRoundPeg
func (r *SimpleRoundPeg) GetRadius() float64 {
	return r.Radius
}

// SquarePeg (incompatible with RoundHole)
type SquarePeg struct {
	Width float64
}

// SquarePegAdapter adapts SquarePeg to RoundPeg
type SquarePegAdapter struct {
	SquarePeg *SquarePeg
}

// GetRadius calculates the equivalent radius of the square peg
func (s *SquarePegAdapter) GetRadius() float64 {
	return s.SquarePeg.Width * math.Sqrt2 / 2
}

func main() {
	hole := &RoundHole{Radius: 5}

	roundPeg := &SimpleRoundPeg{Radius: 5}
	fmt.Println("Round peg fits:", hole.Fits(roundPeg)) // ✅ Fits

	squarePeg := &SquarePeg{Width: 7}
	//fmt.Println("Round peg fits:", hole.Fits(squarePeg)) // SquarePeg (incompatible with RoundHole)
	// -> Using adapter to convert squarePeg to roundPeg

	adapter := &SquarePegAdapter{SquarePeg: squarePeg}

	fmt.Println("Square peg fits using adapter:", hole.Fits(adapter)) // Compatible with RoundHole
}

/*
	1 . RoundHole expects objects that implement the RoundPeg interface.
	2 . SquarePeg is incompatible because it doesn't have a GetRadius() method.
	3 . SquarePegAdapter adapts SquarePeg by computing an equivalent radius (width * sqrt(2) / 2).
	Now, we can use a square peg where a round peg is required.

*/
