package go_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// worker池实现
func TestWorkPool(t *testing.T) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	worker := func(id int, jobs <-chan int, results chan<- int) {
		for job := range jobs {
			fmt.Println("worker ", id, "processing job ", job)
			time.Sleep(time.Second)
			results <- job * 2
		}
	}

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 10; i++ {
		result := <-results
		println(result)
	}
}

// 生产者消费者实现
func TestFanInFanOut(t *testing.T) {
	producer := func(ch chan<- int, d time.Duration) {
		var i int
		for {
			ch <- i
			i++
			time.Sleep(d)
		}
	}

	consumer := func(ch <-chan int) {
		for v := range ch {
			fmt.Println(v)
		}
	}

	ch := make(chan int)

	go producer(ch, time.Millisecond*500)
	go producer(ch, time.Millisecond*2000)

	consumer(ch)
}

// 关闭chan
func TestWait(t *testing.T) {
	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	i, ok := <-c
	println(i, ok)
	//配合select使用可以来检测channel的关闭事件
	//select {
	//case <-c:
	//	fmt.Printf("Unblocked %v later.\n", time.Since(start))
	//}
}

func TestContextTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	taskCh := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		taskCh <- 1
	}()

	select {
	case data := <-taskCh:
		fmt.Println(data)
	case <-ctx.Done():
		fmt.Println("timeout!")
	}
}
