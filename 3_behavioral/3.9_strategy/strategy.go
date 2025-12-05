package main

import (
	"fmt"
)

// 策略模式: 定义一些独立的类来封装不同的算法, 每一个类封装一个具体的算法

// 定义一个策略类

type IStrategy interface {
	do(int, int) int
}

type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct{}

func (*reduce) do(a, b int) int {
	return a - b
}

// 具体的策略执行者
type operator struct {
	strategy IStrategy
}

// 设置策略
func (r *operator) setStrategy(s IStrategy) {
	r.strategy = s
}

// 执行策略
func (r *operator) calculate(a, b int) int {
	return r.strategy.do(a, b)
}

func main() {
	opt := &operator{}
	opt.setStrategy(&add{})
	fmt.Println(opt.calculate(1, 2))

	opt.setStrategy(&reduce{})
	fmt.Println(opt.calculate(2, 1))
}
