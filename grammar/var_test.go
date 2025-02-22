package grammar

import (
	"fmt"
	"testing"
)

// 声明变量的一般形式是使用 var 关键字：
// var identifier type
// 可以一次声明多个变量：
// var identifier1, identifier2 type
func TestVarI(t *testing.T) {
	var vInt1, VInt2 int = 1, 2
	fmt.Println(vInt1, VInt2)

	// 声明一个变量并初始化
	var a string = "golang"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

	//字符串为 ""（空字符串）
	//布尔类型为 false
	//数值类型（包括complex64/128）为 0
	//以下几种类型为 nil：
	//var a *int
	//var a []int
	//var a map[string] int
	//var a chan int
	//var a func(string) int
	//var a error

	//Go 语言常量
	//const identifier [type] = value
	const LENGTH int = 10
	const WIDTH int = 5
}
