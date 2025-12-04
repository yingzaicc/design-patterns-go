package main

import (
	"sync"
)

// 懒汉模式: 延迟创建实例

type Singleton struct{}

var (
	instance *Singleton
	mu       sync.Mutex
	once     sync.Once
)

// GetInstance01 简单实现,线程不安全,不建议使用
func GetInstance01() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// GetInstance02 两次非空判断,线程安全
func GetInstance02() *Singleton {
	if instance == nil {
		mu.Lock()
		if instance == nil {
			instance = &Singleton{}
		}
		mu.Unlock()
	}
	return instance
}

// GetInstance03 利用 once.Do 全局只创建一次
func GetInstance03() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	//GetInstance01()
	GetInstance02()
	GetInstance03()
}
