package go_test

import (
	"fmt"
	"testing"
	"time"
)

// select 语句允许一个 goroutine 等待多个通信操作中的任意一个完成。它类似于 switch 语句，但用于 channel 操作。
func TestSelect00(t *testing.T) {
	ch := make(chan int, 1)

	ch <- 1
	//有时候打印random 01、有时候打印random 02;
	//如果多个通道都准备好，那么 select 语句会随机选择一个通道执行
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	}
}

func TestSelect01(t *testing.T) {
	ch := make(chan int, 1)

	//如果所有通道都没有准备好,fatal error: all goroutines are asleep - deadlock!
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	}
}

func TestSelect02(t *testing.T) {
	ch := make(chan int, 1)

	//如果所有通道都没有准备好，那么执行 default 块中的代码。
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	default:
		fmt.Println("random 03")
	}
}

// 超时等待实现
func TestSelect03(t *testing.T) {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	select {
	case data := <-ch:
		fmt.Println(data)
	case <-time.After(time.Second * 5):
		fmt.Println("timeout!")
	}

}

// 检测chan是否以及满
func TestSelect04(t *testing.T) {
	ch := make(chan string, 1)

	// 向缓冲 channel 发送数据
	ch <- "初始消息"

	//goroutine 34 [chan send]:
	//go-test_test.TestSelect(0x14000118820?) 如果没有default，case ch <- "新的消息":是一个阻塞操作
	//select {
	//case ch <- "新的消息":
	//	fmt.Println("发送成功")
	//}

	select {
	case ch <- "新的消息":
		fmt.Println("发送成功")
	default:
		fmt.Println("发送失败，channel 已满")
	}

}

// 将可迭代的value转换为通道上的值
func TestSelect05(t *testing.T) {
	stringStream := make(chan string, 3)
	for _, s := range []string{"a", "b", "c"} {
		select {
		case stringStream <- s:
		}
	}

	go func() {
		for i := 0; i < 3; i++ {
			char := <-stringStream
			fmt.Println(char)
		}
	}()
}

// 无限循环等待停止
func TestSelect06(t *testing.T) {
	done := make(chan bool)

	go func() {
		time.Sleep(time.Second * 3)
		close(done)
	}()

	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("sleep ...")
			time.Sleep(time.Second)
		}
	}

}
