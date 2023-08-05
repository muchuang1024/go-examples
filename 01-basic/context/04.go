package main

import (
	"context"
	"fmt"
	"time"
)

func myFunction(ctx context.Context) {
	select {
	case <-ctx.Done():
		// 接收到取消信号，结束运行
		fmt.Println("Function canceled")
		return
	default:
		// 执行一些操作
		fmt.Println("Running...")
		// 模拟一些耗时操作
		time.Sleep(5 * time.Second)
		fmt.Println("Function completed")
	}
}

func main() {
	// 创建带有超时功能的上下文，设置超时时间为3秒
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 启动 Goroutine 并传入上下文
	go myFunction(ctx)

	// 等待上下文超时（主程序结束导致子程序结束，如果主程序永远不结束就有问题）
	select {
	case <-ctx.Done():
		fmt.Println("Main function canceled")
	case <-time.After(5 * time.Second):
		fmt.Println("Main function completed")
	}
}
