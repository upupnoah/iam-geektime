package structural_patterns

import "fmt"

// 结构型模式：用于解决对象之间的 组合 和 接口 问题

// 代理模式（Proxy Pattern）为其他对象提供一种代理以控制对这个对象的访问

type Seller interface {
	sell(name string)
}

// Station 火车站
type Station struct {
	stock int // 库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中: %s买了一张票，剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("票已售罄")
	}
}

// StationProxy 火车代理点
type StationProxy struct {
	station *Station
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("代理点中: %s买了一张票，剩余：%d \n", name, proxy.station.stock)
	} else {
		fmt.Println("票已售罄")
	}
}
