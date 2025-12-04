package main

import (
	"fmt"
	"time"
)

// 选项模式: 选项模式不是 GoF 23/25 种经典模式，是 Go 开发者为适配语言特性发明的 “实用型场景模式”，属于创建型模式的补充
// 创建一个带有默认值的 struct 变量, 并选择性的修改其中一些参数的值

const (
	defaultTimeout = 10 * time.Second
	defaultCaching = false
)

type Connection struct {
	addr    string
	timeout time.Duration
	caching bool
}

// 选项类
type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

// 选项函数

func WithTimeout(timeout time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = timeout
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

// 创建选项

func NewConnect(addr string, opts ...Option) (*Connection, error) {
	option := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o.apply(&option)
	}

	return &Connection{
		addr:    addr,
		timeout: option.timeout,
		caching: option.caching,
	}, nil
}

func main() {
	conn, err := NewConnect("192.168.0.1",
		WithTimeout(time.Microsecond*2),
		WithCaching(true))
	if err != nil {
		panic(err)
	}
	fmt.Printf("conn: %+v\n", conn)
	// output: conn: &{addr:192.168.0.1 timeout:2000 caching:true}
}
