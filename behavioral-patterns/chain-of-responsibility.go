package main

import "fmt"

// Rerfer to: https://refactoring.guru/design-patterns/chain-of-responsibility
/*
Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request,
each handler decides either to process the request or to pass it to the next handler in the chain.
*/
type Handler interface {
	SetNextHandler(handler Handler)
	Handle(string)
}

type zHandler struct {
	nextHandler Handler
}

func (z *zHandler) SetNextHandler(h Handler) {
	z.nextHandler = h
}
func (z *zHandler) HandleNext(lv string) {
	z.nextHandler.Handle(lv)
}

type P1 struct {
	zHandler
}

func (p *P1) Handle(lv string) {
	if lv != "p1" {
		fmt.Println("P1 cannot handle for this")
		p.nextHandler.Handle(lv)
	} else {
		fmt.Println("P1 handle for this")
	}
}

type P2 struct {
	zHandler
}

func (p *P2) Handle(lv string) {
	if lv != "p2" {
		fmt.Println("P2 cannot handle for this")
		p.nextHandler.Handle(lv)
	} else {
		fmt.Println("P2 handle for this")
	}
}

type P3 struct {
	zHandler
}

func (p *P3) Handle(lv string) {
	if lv != "p3" {
		fmt.Println("P3 cannot handle for this")
		p.nextHandler.Handle(lv)
	} else {
		fmt.Println("P3 handle for this")
	}
}

func main() {
	p1 := &P1{}
	p2 := &P2{}
	p3 := &P3{}
	p1.SetNextHandler(p2)
	p2.SetNextHandler(p3)
	p1.Handle("p3")
}
