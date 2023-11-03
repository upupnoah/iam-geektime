package pflag

import "github.com/spf13/pflag"

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
