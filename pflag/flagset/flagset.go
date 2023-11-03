package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

// flagSet的两种定义方式
func createFlagSet() {
	// 使用 pflag 包来解析命令行参数

	// FlagSet 是一些预先定义好的 Flag 的集合，几乎所有的 Pflag 操作，都需要借助 FlagSet 提供的方法来完成
	// FlagSet 有两种创建方式
	// 方法 1: 调用 NewFlagSet 创建一个 FlagSet对象
	var version bool
	fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
	fs.BoolVar(&version, "version", true, "Print version information and quit")

	// 方法 2: 调用 pflag 包的全局函数来定义命令行参数
	// 在一些不需要定义子命令的命令行工具中，我们可以直接使用全局的 FlagSet，更加简单方便
	pflag.BoolVarP(&version, "version", "v", true, "Print version information and quit")
}

func main() {
	// 创建 FlagSet
	var fs = pflag.NewFlagSet("example", pflag.ExitOnError)

	// 定义命令行参数
	var username = fs.String("username", "guest", "Your username")
	var password = fs.StringP("password", "p", "12345", "Your password")

	// 解析命令行参数
	err := fs.Parse(os.Args[1:])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 访问命令行参数的值
	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)

	// 获取非选项参数
	fmt.Println("Non-option arguments:", fs.Args())
}
