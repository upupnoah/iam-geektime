package main

import "fmt"

// Person 在工厂方法模式中，依赖工厂函数，我们可以通过实现工厂函数来创建多种工厂
type Person struct {
	name string
	age  int
}

func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			name: name,
			age:  age,
		}
	}
}

func main() {
	newBaby := NewPersonFactory(1)
	baby := newBaby("john")

	NewTeenager := NewPersonFactory(16)
	teenager := NewTeenager("jill")

	fmt.Println(baby, teenager)
}
