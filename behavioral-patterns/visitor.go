package main

import "fmt"

// Refer to: https://refactoring.guru/design-patterns/visitor
/*
Visitor is a behavioral design pattern that lets you separate
algorithms from the objects on which they operate.
*/

type Cake interface {
	getType() string
	accept(visitor)
}
type aCake struct {
	t string
}

func (aCake *aCake) getType() string {
	return aCake.t
}
func (aCake *aCake) accept(v visitor) {
	v.visitACake(aCake)
}

type bCake struct {
	t string
}

func (bCake *bCake) getType() string {
	return bCake.t
}
func (bCake *bCake) accept(v visitor) {
	v.visitBCake(bCake)
}

// If we want to add more method to existed struct, but we don't want to modify more those struct (Current getType() string method)
// We use Visitor pattern with visitor interface and add only one method accept to those existed struct
type visitor interface {
	visitACake(*aCake)
	visitBCake(*bCake)
}
type nVisitor1 struct{}

func (n1 *nVisitor1) visitACake(c *aCake) {
	fmt.Println("nVisitor1 called", c.getType())
}
func (n1 *nVisitor1) visitBCake(c *bCake) {
	fmt.Println("nVisitor1 called", c.getType())
}

type nVisitor2 struct{}

func (n2 *nVisitor2) visitACake(c *aCake) {
	fmt.Println("nVisitor2 called", c.getType())
}
func (n2 *nVisitor2) visitBCake(c *bCake) {
	fmt.Println("nVisitor2 called", c.getType())
}
func main() {
	a := &aCake{t: "aCakeType"}
	b := &bCake{t: "bCakeType"}
	n1 := &nVisitor1{}
	a.accept(n1)
	b.accept(n1)

	n2 := &nVisitor2{}
	a.accept(n2)
	b.accept(n2)

}
