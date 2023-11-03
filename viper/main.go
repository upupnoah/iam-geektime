package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// 读入方式

// 设置默认的配置文件名
// 设置默认值通常是很有用的，可以让程序在没有明确指定配置时也能够正常运行
func setDefaultConfigName() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}

// 读取配置文件
var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func readConfigFile() {
	// 从配置文件中读取配置
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	} else {
		viper.AddConfigPath(".")          // 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("$HOME/.iam") // 配置文件搜索路径，可以设置多个配置文件搜索路径
		viper.SetConfigName("config")     // 配置文件名称（没有文件扩展名）
	}

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
}

// 监听和重新读取配置文件
func watchConfigFile() {
	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

// 设置配置值
func setConfigValue() {
	viper.Set("user.username", "noah")
}

// 使用环境变量
func useEnv() {
	// 使用环境变量
	os.Setenv("VIPER_USER_SECRET_ID", "QLdywI2MrmDVjSSv6e95weNRvmteRjfKAuNV")
	os.Setenv("VIPER_USER_SECRET_KEY", "bVix2WBv0VPfrDrvlLWrhEdzjLpPCNYb")

	viper.AutomaticEnv()                                             // 读取环境变量
	viper.SetEnvPrefix("VIPER")                                      // 设置环境变量前缀：VIPER_，如果是viper，将自动转变为大写。
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")) // 将viper.Get(key) key字符串中'.'和'-'替换为'_'
	viper.BindEnv("user.secret-key")
	viper.BindEnv("user.secret-id", "USER_SECRET_ID") // 绑定环境变量名到key
}

// 使用标志
func useFlag() {
	viper.BindPFlag("token", pflag.Lookup("token")) // 绑定单个标志
	viper.BindPFlags(pflag.CommandLine)             //绑定标志集
}

// 读取配置
func readConfig() {
	//{
	//	"host": {
	//	"address": "localhost",
	//		"port": 5799
	//},
	//	"datastore": {
	//	"metric": {
	//		"host": "127.0.0.1",
	//			"port": 3099
	//	},
	//	"warehouse": {
	//		"host": "198.0.0.1",
	//			"port": 2112
	//	}
	//}
	//}
	viper.GetString("datastore.metric.host") // (返回 "127.0.0.1")
}

// 反序列化
func unmarshal() {
	type config struct {
		Port    int
		Name    string
		PathMap string `mapstructure:"path_map"`
	}

	var C config

	err := viper.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
}

// 序列化成字符串
func marshal() string {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		log.Fatalf("unable to marshal config to YAML: %v", err)
	}
	return string(bs)
}
func main() {
}
