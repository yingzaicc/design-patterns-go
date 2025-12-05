package main

import (
	"fmt"
)

// 模板方法: 将一个类中能够公共使用的方法放置在抽象类中实现, 将不能公共使用的方法作为抽象方法,
// 强制子类去实现

type Cooker interface {
	fire()
	cooke()
	outfire()
}

// 类似一个抽象类

type CookMenu struct{}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 交给具体的子类实现
func (CookMenu) cooke() {
}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

// 封装具体步骤
func doCook(cooker Cooker) {
	cooker.fire()
	cooker.cooke()
	cooker.outfire()
}

// 具体的子类

type XiHongShi struct {
	CookMenu
}

func (XiHongShi) cooke() {
	fmt.Println("炒西红柿")
}

type JiDan struct {
	CookMenu
}

func (JiDan) cooke() {
	fmt.Println("炒鸡蛋")
}

func main() {
	doCook(XiHongShi{})
	doCook(JiDan{})
}
