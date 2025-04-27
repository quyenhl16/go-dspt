package main

import "fmt"

// Reference: https://refactoring.guru/design-patterns/factory-method
/* Factory Method is a creational design pattern that provides an interface for creating objects in a superclass,
but allows subclasses to alter the type of objects that will be created.
*/

type IFactory interface {
	Cake()
	Milk()
}
type CakeFactory struct {
	name string
}

// NewCake returns a new CakeFactory instance
// But NewCake function returns IFactory instead of *CakeFactory
func NewCake() IFactory {
	return &CakeFactory{
		name: "BongLan",
	}
}

func (c *CakeFactory) Cake() {
	fmt.Println("Cake factory produces cakes")
}
func (c *CakeFactory) Milk() {
	fmt.Println("Cake factory does not produce milk")
}

type MilkFactory struct {
	name string
}

// NewMilk return a new MilkFactory instance
// But NewMilk function returns IFactory instead of *MilkFactory
func NewMilk() IFactory {
	return &MilkFactory{
		name: "TH TrueMilk",
	}
}

func (c *MilkFactory) Cake() {
	fmt.Println("Milk factory does not produce cakes")
}
func (c *MilkFactory) Milk() {
	fmt.Println("Milk factory produces milk")
}

func getFactory(nameF string) IFactory {
	if nameF == "cake" {
		return NewCake()
	} else {
		return NewMilk()
	}
}

func main() {
	var manyFactorys = []IFactory{NewCake(), NewMilk(), NewMilk()} // (1)
	for _, f := range manyFactorys {
		f.Cake()
		f.Milk()
	}

}

/* Conclusion: So if we have more factory like: clothes, household, ... We just add corresponding struct and method
and no have modified function getFactory refer than We return *Factory
We can add more factory to IFactory slice
*/
