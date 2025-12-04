package main

import (
	"fmt"
)

// 简单工厂: 接收参数, 返回实例

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("name: %s, age: %d\n", p.Name, p.Age)
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func main() {
	p := NewPerson("张三", 18)
	p.Greet()
}
