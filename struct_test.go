package go_test

import (
	"fmt"
	"testing"
)

type MyStruct struct {
	Num int
}

//在 Go 语言中，为同一个结构体类型（例如 MyStruct）同时定义值接收者（value receiver）和指针接收者（pointer receiver）的做法是不推荐的。
//为什么不推荐混用值接收者和指针接收者？
//一致性和可预测性：一致的行为：如果一个类型的方法集合中既有值接收者又有指针接收者，使用该类型的代码在调用这些方法时可能会有不同的行为。例如，某些方法可能会修改接收者，而另一些则不会，这可能会让使用者感到困惑。

// 指针接收器，指针方法
func (mi *MyStruct) print() {
	fmt.Println("MyStruct:", (*mi).Num)
}

// 值接收器，值方法
func (mi MyStruct) echo() {
	fmt.Println("MyStruct:", mi.Num)
}
func Test2(t *testing.T) {
	i := &MyStruct{0}
	i.print()
	i.echo()
	//(*i).echo()

	var j MyStruct = MyStruct{6}
	//(&j).print()
	j.print()
	j.echo()
}
