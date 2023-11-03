package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

// Pflag使用方法
func commandlineParameterParse() {
	// Pflag使用方法
	// 支持长选项、默认值和使用文本，并将标志的值存储在指针中
	var _ = pflag.String("name", "colin", "Input Your Name")
	// 比上面的多了一个短选项
	var _ = pflag.StringP("name", "n", "colin", "Input Your Name")

	//支持长选项、默认值和使用文本，并将标志的值绑定到变量
	var name string
	pflag.StringVar(&name, "name", "colin", "Input Your Name")
	// 比上面的多了一个短选项
	var name1 string
	pflag.StringVarP(&name1, "name", "n", "colin", "Input Your Name")
}

// 使用Get<Type>获取参数的值
func useGetType() {
	var version bool
	flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	flagSet.BoolVar(&version, "version", true, "Print version information and quit")
	_, _ = flagSet.GetInt("flagname")
}

// 获取非选项参数
func getNonOptionArgs() {
	_ = pflag.Int("flagname", 1234, "help message for flagname")
	pflag.Parse()
	fmt.Printf("argument number is: %v\n", pflag.NArg())
	fmt.Printf("argument list is: %v\n", pflag.Args())
	fmt.Printf("the first argument is: %v\n", pflag.Arg(0))
}

func defaultValue() {
	// 使用 flagset
	//ip := pflag.IntP("flagname", "f", 1234, "help message")
	fg := pflag.NewFlagSet("test", pflag.ContinueOnError)
	//ip := fg.IntP("flagname", "f", 1234, "help message")
	_ = fg.Parse([]string{"--flagname=148", "a", "b", "c"}) // 获取的是=后面的值, 并且不包括非选项参数
	//fg.Lookup("flagname").NoOptDefVal = "4321"

	// 使用全局函数
	ip := pflag.IntP("flagname", "f", 1234, "help message")
	pflag.Parse()
	pflag.Lookup("flagname").NoOptDefVal = "4321"
	fmt.Println(*ip)

	// output
	// [nothing]  ip = 1234
	// --flagname=148  ip = 148
	// --flagname ip = 4321
}

// 弃用标志 或者 标志简写
func deprecatedFlag() {
	// mark a flag as deprecated
	pflag.StringP("flagname", "f", "default", "help message for flagname")
	err := pflag.CommandLine.MarkDeprecated("flagname", "please use --flagname instead")
	if err != nil {
		return
	}
	pflag.Parse()
	// go run commandline_parameter_parse.go -f 123
	// go run commandline_parameter_parse.go --f=123
	// go run commandline_parameter_parse.go --flagname=123
}

// 弃用标志简写, 保留标志
func deprecatedShortFlag() {
	// mark a flag shorthand as deprecated
	pflag.StringP("port", "p", "default", "help message for flagname")
	err := pflag.CommandLine.MarkShorthandDeprecated("port", "please use --port instead")
	if err != nil {
		return
	}
	pflag.Parse()
	// go run commandline_parameter_parse.go -p 123
	// go run commandline_parameter_parse.go -p=123
	// go run commandline_parameter_parse.go --port=123
}

// 隐藏标志
func hiddenFlag() {
	// hide a flag by specifying its name
	//pflag.StringP("secretFlag", "s", "default", "help message for flagname")

	// 绑定到变量中
	var secretFlag string
	pflag.StringVarP(&secretFlag, "secretFlag", "s", "default", "help message for flagname")
	err := pflag.CommandLine.MarkHidden("secretFlag")
	if err != nil {
		return
	}
	pflag.Parse()
	// go run commandline_parameter_parse.go -s 123
	// go run commandline_parameter_parse.go -s=123
	// go run commandline_parameter_parse.go --secretFlag=123
	fmt.Println(secretFlag)
}

func main() {
	//fg := pflag.NewFlagSet("test", pflag.ContinueOnError)
	//fg.Int("flagname", 1234, "help message for flagname")
	//_ = fg.Parse([]string{"--flagname=2345", "a", "b", "c"}) // 获取的是=后面的值, 并且不包括非选项参数
	//
	//i, err := fg.GetInt("flagname")
	//
	//if err != nil {
	//	panic(err)
	//}
	//println(i)

	//getNonOptionArgs()

	//defaultValue()

	//deprecatedFlag()

	//deprecatedShortFlag()

	hiddenFlag()
}
