package singleton

import "sync"

// 懒汉：在第一次使用到的地方创建

// Singleton2 非并发安全版本
type Singleton2 struct{}

var ins2 *Singleton2

func GetInsOr2() *Singleton2 {
	if ins == nil {
		ins2 = &Singleton2{}
	}
	return ins2
}

// Singleton3 并发安全版本（加锁）
type Singleton3 struct{}

var ins3 *Singleton3
var mu sync.Mutex

func GetIns() *Singleton3 {
	if ins3 == nil {
		mu.Lock()
		if ins == nil {
			ins3 = &Singleton3{}
		}
		mu.Unlock() // 这里不用 defer 是因为，defer 会在函数返回之前调用，会影响并发性能
	}
	return ins3
}

// Singleton4 并发安全版本（Go 优雅写法）
type Singleton4 struct{}

var ins4 *Singleton4
var once sync.Once

func GetInsOr4() *Singleton4 {
	once.Do(func() {
		ins4 = &Singleton4{}
	})
	return ins4
}
