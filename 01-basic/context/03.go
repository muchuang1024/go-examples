package main

import (
	"context"
	"fmt"
	"time"
)

func myFunction(ctx context.Context) {
	// 必须是for循环
	for {
		select {
		case <-ctx.Done():
			// 接收到取消信号，结束 Goroutine 运行
			fmt.Println("Goroutine canceled")
			return
		default:
			// 执行一些操作
			fmt.Println("Running...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// 创建带有取消功能的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 Goroutine 并传入上下文
	go myFunction(ctx)

	// 通过Time.After实现取消 Goroutine 运行
	time.AfterFunc(3*time.Second, cancel)

	// 等待一段时间，以确保 Goroutine 结束
	time.Sleep(1 * time.Second)

	fmt.Println("Main function ended")
}
