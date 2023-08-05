package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	goroutineCount := 3

	// 告诉 signal 包要监听哪些信号
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// 启动多个 Goroutine 接收信号
	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cond.L.Lock()
			cond.Wait()
			cond.L.Unlock()
			for sig := range signalChan {
				fmt.Printf("Goroutine %d 收到信号: %v\n", id, sig)
			}
		}(i)
	}

	// 发送终止信号给所有 Goroutine
	fmt.Println("发送终止信号给所有 Goroutine")
	cond.Broadcast()

	// 等待所有 Goroutine 结束
	fmt.Println("等待所有 Goroutine 结束")
	wg.Wait()
}
