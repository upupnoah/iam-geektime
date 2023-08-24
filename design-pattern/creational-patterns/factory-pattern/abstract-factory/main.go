package abstract_factory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

// 抽象工厂模式与简单工厂模式的区别：
//     返回接口，而不是结构体

type Person interface {
	Greet()
}

type person struct {
	name string
	age  int
}

func (p person) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// NewPerson returns an interface, and not the person struct itself
func NewPerson(name string, age int) Person {
	return person{
		name: name,
		age:  age,
	}
}

// 实现多个工厂函数，返回不同的接口实现
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewHttpClient gives us a regular HTTP client from the `net/http` package
func NewHttpClient() Doer {
	return &http.Client{}
}

type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// The `NewRecorder` method of the httptest package gives us
	// a new mock request generator
	res := httptest.NewRecorder()

	// calling the `Result` method gives us
	// the default empty *http.Response object
	return res.Result(), nil
}

// NewMockHTTPClient gives us a mock HTTP client, which returns
// an empty response for any request sent to it
func NewMockHTTPClient() Doer {
	return &mockHTTPClient{}
}
