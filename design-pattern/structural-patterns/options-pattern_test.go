package structural_patterns

import (
	"fmt"
	"testing"
	"time"
)

func TestOptionsPattern(t *testing.T) {
	// 使用默认选项创建连接对象
	conn1, err := NewConnect3("example.com")
	if err != nil {
		fmt.Println("Failed to create connection:", err)
		return
	}
	fmt.Println("Connection 1:", conn1)

	// 使用自定义选项创建连接对象
	conn2, err := NewConnect3("example.com", WithTimeout(5*time.Second), WithCaching(true))
	if err != nil {
		fmt.Println("Failed to create connection:", err)
		return
	}
	fmt.Println("Connection 2:", conn2)
}
