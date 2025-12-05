package main

import (
	"fmt"
)

// 抽象工厂: 返回接口, 而不是结构体

type Animal interface {
	Greet()
}

type Cat struct {
	Name string
}

func (p Cat) Greet() {
	fmt.Printf("Hi, Cat name: %s\n", p.Name)
}

type Dog struct {
	Name string
}

func (p Dog) Greet() {
	fmt.Printf("Hi, Dog name: %s\n", p.Name)
}

func NewCatFactory(name string) Animal {
	return Cat{
		Name: name,
	}
}

func NewDogFactory(name string) Animal {
	return Dog{
		Name: name,
	}
}

func main() {
	NewCatFactory("Tom").Greet()
	NewDogFactory("Jerry").Greet()
}
