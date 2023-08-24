package behavioral_patterns

import "fmt"

// 模版模式（Template Pattern）定义一个操作中的算法的骨架，而将一些步骤延迟到子类中

type Cooker interface {
	fire()
	cooke()
	outfire()
}

// CookMenu 类似于一个抽象类
type CookMenu struct{}

func (CookMenu) fire() {
	fmt.Println("开火")
}

// 做菜，交给具体的子类实现
func (CookMenu) cooke() {}

func (CookMenu) outfire() {
	fmt.Println("关火")
}

// 封装具体的步骤
func doCook(cook Cooker) {
	cook.fire()
	cook.cooke()
	cook.outfire()
}

type XiHongShi struct {
	CookMenu
}

func (*XiHongShi) cooke() {
	fmt.Println("做西红柿")
}

type ChaoJiDan struct {
	CookMenu
}

func (*ChaoJiDan) cooke() {
	fmt.Println("做炒鸡蛋")
}
