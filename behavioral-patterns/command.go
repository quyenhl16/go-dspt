package main

import "fmt"

// Shape is the base interface for all drawable shapes.
// It defines the core contract that all concrete shapes must implement.
// This interface supports both the Visitor pattern (via Accept method) and
// the Prototype pattern (via Clone method).
type Shape interface {
	Accept(Visitor)  // Accepts a visitor - enables operations to be performed on shapes without modifying them
	Clone() Shape    // Clone for prototype pattern - allows creating copies of existing shapes
	GetName() string // Returns the name of the shape for identification
}

// Visitor interface defines operations that can be performed on different shapes.
// This is a core component of the Visitor pattern which allows adding new operations
// to existing object structures without modifying them.
type Visitor interface {
	VisitCircle(*Circle)       // Method to visit Circle objects
	VisitRectangle(*Rectangle) // Method to visit Rectangle objects
}

// Circle implements the Shape interface and represents a circular shape.
// It contains properties specific to circles.
type Circle struct {
	Radius int    // The radius of the circle
	Color  string // The color of the circle
}

// Accept implements the Shape interface for Circle.
// It calls the appropriate visitor method for Circles.
// This is part of the double-dispatch mechanism in the Visitor pattern.
func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

// Clone implements the Shape interface for Circle.
// Returns a deep copy of the Circle object.
// This is part of the Prototype pattern implementation.
func (c *Circle) Clone() Shape {
	return &Circle{Radius: c.Radius, Color: c.Color}
}

// GetName returns the type name of this shape.
// Used for identification in command execution and other operations.
func (c *Circle) GetName() string {
	return "Circle"
}

// Rectangle implements the Shape interface and represents a rectangular shape.
// It contains properties specific to rectangles.
type Rectangle struct {
	Width  int    // The width of the rectangle
	Height int    // The height of the rectangle
	Color  string // The color of the rectangle
}

// Accept implements the Shape interface for Rectangle.
// It calls the appropriate visitor method for Rectangles.
// This is part of the double-dispatch mechanism in the Visitor pattern.
func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

// Clone implements the Shape interface for Rectangle.
// Returns a deep copy of the Rectangle object.
// This is part of the Prototype pattern implementation.
func (r *Rectangle) Clone() Shape {
	return &Rectangle{Width: r.Width, Height: r.Height, Color: r.Color}
}

// GetName returns the type name of this shape.
// Used for identification in command execution and other operations.
func (r *Rectangle) GetName() string {
	return "Rectangle"
}

// Command interface defines the contract for all executable commands.
// This is the core of the Command pattern which encapsulates requests as objects.
type Command interface {
	Execute()       // Executes the command action
	Clone() Command // Creates a copy of the command for history tracking
}

// DrawCommand implements the Command interface for drawing shapes.
// It encapsulates the action of drawing a shape as a command object.
type DrawCommand struct {
	shape Shape // The shape to be drawn
}

// Execute implements the Command interface for DrawCommand.
// When executed, it draws the associated shape.
func (c *DrawCommand) Execute() {
	fmt.Println("Drawing", c.shape.GetName())
}

// Clone implements the Command interface for DrawCommand.
// Creates a copy of the command, ensuring the shape is also cloned.
// This demonstrates how Prototype pattern is used with Command pattern.
func (c *DrawCommand) Clone() Command {
	// Clone the shape too to ensure a deep copy
	return &DrawCommand{shape: c.shape.Clone()}
}

// MoveCommand implements the Command interface for moving shapes.
// It encapsulates the action of moving a shape as a command object.
type MoveCommand struct {
	shape  Shape // The shape to be moved
	dx, dy int   // The distance to move in x and y directions
}

// Execute implements the Command interface for MoveCommand.
// When executed, it moves the associated shape by the specified distance.
func (c *MoveCommand) Execute() {
	fmt.Printf("Moving %s by dx=%d, dy=%d\n", c.shape.GetName(), c.dx, c.dy)
}

// Clone implements the Command interface for MoveCommand.
// Creates a copy of the command, ensuring the shape is also cloned.
// This demonstrates how Prototype pattern is used with Command pattern.
func (c *MoveCommand) Clone() Command {
	return &MoveCommand{shape: c.shape.Clone(), dx: c.dx, dy: c.dy}
}

// CommandInvoker manages and executes commands.
// It also keeps a history of executed commands for potential undo operations.
// This is part of the Command pattern implementation.
type CommandInvoker struct {
	history []Command // History of executed commands
}

// ExecuteCommand executes a command and stores a clone in history.
// By storing clones rather than references to the original commands,
// we ensure that the history is independent of future changes to the commands.
func (i *CommandInvoker) ExecuteCommand(cmd Command) {
	// Save clone to history for undo/redo capabilities
	i.history = append(i.history, cmd.Clone())
	cmd.Execute()
}

// ShowHistory displays the command execution history.
// This is useful for debugging and understanding the sequence of operations.
func (i *CommandInvoker) ShowHistory() {
	fmt.Println("\nCommand history:")
	for i, cmd := range i.history {
		fmt.Printf("%d: %+v\n", i+1, cmd)
	}
}

// RenderVisitor implements the Visitor interface for rendering shapes.
// It defines how different shapes should be rendered to the screen.
// This demonstrates the Visitor pattern implementation.
type RenderVisitor struct{}

// VisitCircle implements the Visitor interface for Circle objects.
// It defines how to render a Circle.
func (v *RenderVisitor) VisitCircle(c *Circle) {
	fmt.Printf("Rendering a %s Circle of radius %d\n", c.Color, c.Radius)
}

// VisitRectangle implements the Visitor interface for Rectangle objects.
// It defines how to render a Rectangle.
func (v *RenderVisitor) VisitRectangle(r *Rectangle) {
	fmt.Printf("Rendering a %s Rectangle of %dx%d\n", r.Color, r.Width, r.Height)
}

// ExportVisitor implements the Visitor interface for exporting shapes to SVG.
// It defines how different shapes should be exported to SVG format.
// This demonstrates how the Visitor pattern allows adding new operations
// without modifying the shape classes.
type ExportVisitor struct{}

// VisitCircle implements the Visitor interface for Circle objects.
// It defines how to export a Circle to SVG format.
func (v *ExportVisitor) VisitCircle(c *Circle) {
	fmt.Printf("Exporting Circle to SVG: <circle r=\"%d\" fill=\"%s\" />\n", c.Radius, c.Color)
}

// VisitRectangle implements the Visitor interface for Rectangle objects.
// It defines how to export a Rectangle to SVG format.
func (v *ExportVisitor) VisitRectangle(r *Rectangle) {
	fmt.Printf("Exporting Rectangle to SVG: <rect width=\"%d\" height=\"%d\" fill=\"%s\" />\n", r.Width, r.Height, r.Color)
}

func main() {
	// Create concrete shape instances
	circle := &Circle{Radius: 10, Color: "red"}
	rect := &Rectangle{Width: 20, Height: 15, Color: "blue"}

	// Create a command invoker to manage commands
	invoker := &CommandInvoker{}

	// Create command objects that encapsulate operations on shapes
	drawCircleCmd := &DrawCommand{shape: circle}
	moveCircleCmd := &MoveCommand{shape: circle, dx: 5, dy: 3}
	drawRectCmd := &DrawCommand{shape: rect}

	// Execute commands through the invoker
	// This demonstrates the Command pattern in action
	invoker.ExecuteCommand(drawCircleCmd)
	invoker.ExecuteCommand(moveCircleCmd)
	invoker.ExecuteCommand(drawRectCmd)

	// View command history - demonstrates how the Prototype pattern
	// helps maintain a proper history of commands for potential undo operations
	invoker.ShowHistory()

	// Demonstrate the Visitor pattern in action
	fmt.Println("\n--- Visitors in action ---")
	renderVisitor := &RenderVisitor{}
	exportVisitor := &ExportVisitor{}

	// Each shape accepts different visitors, which perform different operations
	// This demonstrates the power of the Visitor pattern - adding new operations
	// without modifying the shape classes
	circle.Accept(renderVisitor)
	circle.Accept(exportVisitor)
	rect.Accept(renderVisitor)
	rect.Accept(exportVisitor)
}
