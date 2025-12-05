package main

// 饿汉模式: 提前创建实例
type Singleton struct{}

var instance *Singleton = &Singleton{}

func GetInstance() *Singleton {
	return instance
}

func main() {
	_ = GetInstance()
}
