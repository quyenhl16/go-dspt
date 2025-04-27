package main

import (
	"fmt"
	"math/rand"
)

// Reference: https://refactoring.guru/design-patterns/abstract-factory
/* Abstract Factory is a creational design pattern that lets you produce families of
related objects without specifying their concrete classes.
*/
/* Cake factory */
type iCakeFactory interface {
	name() string
	expired() bool
}
type aCake struct {
	n   string
	exp bool
}
type bCake struct {
	n   string
	exp bool
}

func (r *aCake) name() string {
	return r.n
}
func (r *aCake) expired() bool {
	return r.exp
}
func (r *bCake) name() string {
	return r.n
}
func (r *bCake) expired() bool {
	return r.exp
}

/* Milk Factory */
type iMilkFactory interface {
	name() string
	cool() bool
	expired() bool
}
type aMilk struct {
	n   string
	c   bool
	exp bool
}
type bMilk struct {
	n   string
	c   bool
	exp bool
}

func (r *aMilk) name() string {
	return r.n
}
func (r *aMilk) cool() bool {
	return r.c
}
func (r *aMilk) expired() bool {
	return r.exp
}

func (r *bMilk) name() string {
	return r.n
}
func (r *bMilk) cool() bool {
	return r.c
}
func (r *bMilk) expired() bool {
	return r.exp
}

/* Abstract Factory Interface */

type iFoodDrinkFactory interface {
	makeCake() iCakeFactory
	makeMilk() iMilkFactory
}
type aBrand struct {
	bName string
}

func (r *aBrand) makeCake() iCakeFactory {
	return &aCake{n: "Cake Of A"}
}
func (r *aBrand) makeMilk() iMilkFactory {
	return &aMilk{n: "Milk Of A"}
}

func (r *bBrand) makeCake() iCakeFactory {
	return &bCake{n: "Cake Of B"}
}
func (r *bBrand) makeMilk() iMilkFactory {
	return &bMilk{n: "Milk Of B"}
}

type bBrand struct {
	bName string
}

func getFDFactory(brand string) iFoodDrinkFactory {
	switch brand {
	case "A":
		return &aBrand{bName: "This is A brand"}
	case "B":
		return &bBrand{bName: "This is B brand"}
	}
	return nil
}

func main() {
	brand := func() string {
		a := rand.Intn(100)
		if a%2 == 0 {
			return "A"
		}
		return "B"
	}
	f := getFDFactory(brand())
	switch f.(type) {
	case *aBrand:
		fmt.Printf("Brand: %s \n", f.(*aBrand).bName)
	default:
		fmt.Printf("Brand: %s \n", f.(*bBrand).bName)

	}
	fmt.Printf("CakeName: %s \n", f.makeCake().name())
	fmt.Printf("MilkName: %s \n", f.makeMilk().name())

	newBrand := "B"
	nf := getFDFactory(newBrand)
	switch nf.(type) {
	case *aBrand:
		fmt.Printf("Brand: %s \n", nf.(*aBrand).bName)
	default:
		fmt.Printf("Brand: %s \n", nf.(*bBrand).bName)

	}

	fmt.Printf("CakeName: %s \n", nf.makeCake().name())
	fmt.Printf("MilkName: %s \n", nf.makeMilk().name())

}
