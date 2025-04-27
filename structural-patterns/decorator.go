// Refer to: https://refactoring.guru/design-patterns/decorator

/*
	Decorator is a structural design pattern that lets you attach new behaviors to objects

by placing these objects inside special wrapper objects that contain the behaviors.
*/
package main

import "fmt"

type IPizza interface {
	getPrice() int
}
type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 10
}

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 100
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 1000
}

type saladToping struct {
	pizza IPizza
}

func (s *saladToping) getPrice() int {
	pizzaPrice := s.pizza.getPrice()
	return pizzaPrice + 10000
}

func main() {
	normalPizza := &VeggieMania{}
	withCheesePizza := &CheeseTopping{pizza: normalPizza}
	withTomato := &TomatoTopping{pizza: withCheesePizza}
	withSalad := &saladToping{pizza: withTomato}
	fmt.Println(withSalad.getPrice())
}
