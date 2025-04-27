package main

import (
	"fmt"
	"sync"
)

// Reference: https://refactoring.guru/design-patterns/singleton

/* Singleton is a creational design pattern that lets you ensure that a class has only one instance,
   while providing a global access point to this instance
*/

// Implement 1: Using func init() in Go
type singletonStruct struct {
	value int
}

var instance1 *singletonStruct // Create global instance

func init() {
	instance1 = &singletonStruct{value: 100}
}

// getInstance1 returns global instance
func getInstance1() *singletonStruct {
	return instance1
}

// Implement 2: Using sync.Once
var instance2 *singletonStruct
var one sync.Once

// getInstance2 returns global instance2
func getInstance2() *singletonStruct {
	if instance2 == nil {
		one.Do(func() {
			instance2 = &singletonStruct{value: 200}
		})
	} else {
		fmt.Println("Instance2 not nill")
	}
	return instance2
}
func main() {
	i1 := getInstance1()
	fmt.Printf("i1: %p\n", i1)

	i2 := getInstance2()
	fmt.Printf("i2: %p\n", i2)
	i3 := getInstance2()
	fmt.Printf("i3: %p\n", i3)
	i4 := getInstance2()
	fmt.Printf("i3: %p\n", i4)
	/* i2, i3, i4 have same address with instance2. */
}
