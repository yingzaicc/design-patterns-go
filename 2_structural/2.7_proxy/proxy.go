package main

import (
	"fmt"
)

// 代理模式: 为另一个对象提供一个替身或占位符, 以控制对这个对象的访问

type Seller interface {
	sell()
}

// 火车站

type Station struct {
	stock int // 库存
}

func (r *Station) sell(name string) {
	if r.stock > 0 {
		r.stock--
		fmt.Printf("火车站: %s 成功购票\n", name)
	} else {
		fmt.Printf("火车站: 已售罄")
	}
}

// 火车站代理点

type StationProxy struct {
	station *Station
}

func (r *StationProxy) sell(name string) {
	if r.station.stock > 0 {
		r.station.sell(name)
	} else {
		fmt.Printf("火车站代理点: 已售罄")
	}
}

func main() {
	station := Station{stock: 1}
	stationProxy := StationProxy{station: &station}
	station.sell("张三")
	stationProxy.sell("李四")
}
