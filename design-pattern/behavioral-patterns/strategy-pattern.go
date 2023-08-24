package behavioral_patterns

// 策略模式（Strategy Pattern）定义一组算法，将每个算法都封装起来，并且使它们之间可以互换
// 行为型：对象之间通信和交互的方式

type IStrategy interface {
	do(int, int) int
}

// 策略实现：加法
type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

// 策略实现：减法
type reduce struct{}

func (*reduce) do(a, b int) int {
	return a - b
}

// Operator 具体策略的执行者
type Operator struct {
	strategy IStrategy
}

// 设置策略
func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// 调用策略中的方法
func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
