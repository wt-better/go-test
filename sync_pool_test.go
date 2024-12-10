package go_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
)

// sync.Pool适用于以下场景：
//
// 高频创建和销毁对象：例如网缓冲区等，频繁创建销毁会增加 GC 压力。
// 短生命周期的临时对象：对象生命周期短且不需要跨多个函数或 goroutine 使用。
// 复用昂贵的对象：如大规模的缓冲区

// 定义对象池
var bufferPool = sync.Pool{
	//这里定义了一个 `bufferPool`，它的 `New` 字段指定了当池中没有 `bytes.Buffer` 对象时，如何创建一个新的 `bytes.Buffer`。
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func process(data string) string {
	// 从池中获取一个 bytes.Buffer
	buf := bufferPool.Get().(*bytes.Buffer)
	//使用完毕后，将对象放回池中以供复用。
	defer bufferPool.Put(buf)

	//在使用对象前，通常需要重置其状态，以确保之前的使用痕迹被清除。
	buf.Reset()

	// 进行一些操作
	buf.WriteString("Processed: ")
	buf.WriteString(data)

	result := buf.String()
	return result
}

func Test(t *testing.T) {
	var wg sync.WaitGroup
	inputs := []string{"data1", "data2", "data3", "data4"}

	for _, input := range inputs {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			fmt.Println(process(d))
		}(input)
	}

	wg.Wait()
}
