package main

import (
	"fmt"
)

// 工厂方法: 使用多个子工厂返回具体结构

type Cat struct {
	Name string
	Age  int
}

func (p Cat) Greet() {
	fmt.Printf("Hi, cat name: %s\n", p.Name)
}

type Dog struct {
	Name string
	Age  int
}

func (p Dog) Greet() {
	fmt.Printf("Hi, dog name: %s\n", p.Name)
}

func NewCatFactory(name string) Cat {
	return Cat{
		Name: name,
	}
}

func NewDogFactory(name string) Dog {
	return Dog{
		Name: name,
	}
}

// 创建具有默认值的工厂

func NewCatFactoryWithDefaultAge(age int) func(name string) Cat {
	return func(name string) Cat {
		return Cat{
			Name: name,
			Age:  age,
		}
	}
}

func main() {
	NewCatFactory("Tom").Greet()
	NewDogFactory("Jerry").Greet()

	newCatFunc := NewCatFactoryWithDefaultAge(10)
	newCatFunc("Tom02").Greet()
}
