package main

import "fmt"

// Refer to: https://refactoring.guru/design-patterns/flyweight
/*
Flyweight is a structural design pattern that lets you fit more
objects into the available amount of RAM by sharing common parts of state between
multiple objects instead of keeping all of the data in each object.
*/

type Doll struct {
	name        string
	material    string
	productTime string
}

type DollFactory struct {
	f map[string]*Doll
}

func NewDollFactory() *DollFactory {
	return &DollFactory{make(map[string]*Doll)}
}
func (f *DollFactory) getDoll(name, meterial, productTime string) *Doll {
	key := fmt.Sprintf("This Doll is %s-%s-%s ", name, meterial, productTime)
	if v, ok := f.f[key]; ok {
		return v
	}
	d := &Doll{name, meterial, productTime}
	f.f[key] = d
	return d
}
func main() {
	f := NewDollFactory()
	d1 := f.getDoll("A", "Plastic", "2025")
	_ = f.getDoll("B", "PlasticZZ", "2026")
	d3 := f.getDoll("A", "Plastic", "2025")
	fmt.Println(f)
	fmt.Printf("%p %p", d3, d1)
	fmt.Println("is d3 == d1", d3 == d1)
}
