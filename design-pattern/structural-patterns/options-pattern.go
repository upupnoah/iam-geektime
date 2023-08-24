package structural_patterns

import "time"

// 选项模式（Options Pattern）也是 Go 项目开发中经常使用到的模式
// 例如，grpc/grpc-go 的NewServer函数，uber-go/zap 包的New函数都用到了选项模式。
// 使用选项模式，我们可以创建一个带有默认值的 struct 变量，并选择性地修改其中一些参数的值

const (
	defaultTimeout = 10
	defaultCaching = false
)

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

// 方法一：分别开发两个用来创建实例的函数，一个创建带默认值的实例，一个可以定制化创建实例

// NewConnect1 creates a connection.
func NewConnect1(addr string) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   defaultCaching,
		timeout: defaultTimeout,
	}, nil
}

// NewConnectWithOptions creates a connection with options.
func NewConnectWithOptions(addr string, cache bool, timeout time.Duration) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   cache,
		timeout: timeout,
	}, nil
}

// *************************************************************************

// 方法二（优雅）：创建一个带默认值的选项，并用该选项创建实例

type ConnectionOptions struct {
	Caching bool
	Timeout time.Duration
}

func NewDefaultOptions() *ConnectionOptions {
	return &ConnectionOptions{
		Caching: defaultCaching,
		Timeout: defaultTimeout,
	}
}

// NewConnect2 creates a connection with options.
func NewConnect2(addr string, opts *ConnectionOptions) (*Connection, error) {
	return &Connection{
		addr:    addr,
		cache:   opts.Caching,
		timeout: opts.Timeout,
	}, nil
}

// *************************************************************************

// 方法 3（更优雅）：使用选项模式创建实例
type options struct {
	caching bool
	timeout time.Duration
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options) // 选项函数

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeout(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func NewConnect3(addr string, opts ...Option) (*Connection, error) {
	option := options{
		caching: defaultCaching,
		timeout: defaultTimeout,
	}
	for _, o := range opts {
		o.apply(&option)
	}

	return &Connection{
		addr:    addr,
		cache:   option.caching,
		timeout: option.timeout,
	}, nil
}
