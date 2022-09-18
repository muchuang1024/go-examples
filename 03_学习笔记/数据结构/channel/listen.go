package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(stop <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("received stop signal")
			return
		default:
		}
	}
}

// 1. 如果多个 goroutine 都监听同一个 channel，那么 channel 上的数据都**可能随机被某一个 goroutine 取走进行消费**
// 2. 如果多个 goroutine 监听同一个 channel，如果这个 channel 被关闭，则所有 goroutine **都能收到退出信号**
func main() {
	stop := make(chan struct{})
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// 每一个 goroutine 都监听同个 stop channel，将可同时收到 stop 信号.
		go worker(stop)
	}

	// 确保所有 goroutine 已经启动.
	time.Sleep(2 * time.Second)
	// stop <- struct{}{}
	close(stop)

	wg.Wait()
}
