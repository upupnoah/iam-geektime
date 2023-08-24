package simple_factory

import "fmt"

// 简单工厂模式（最常用，最简单）
// New 函数

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}
