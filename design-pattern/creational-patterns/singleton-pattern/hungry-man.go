package singleton

// 饿汉：在包被加载的时候创建

type Singleton struct{}

var ins *Singleton = &Singleton{}

func GetInsOr() *Singleton {
	return ins
}
