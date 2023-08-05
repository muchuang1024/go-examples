package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			// 接收到取消信号，结束耗时操作
			fmt.Println("Task canceled")
			return
		default:
			// 执行耗时操作
			fmt.Println("Running task...")
			for {
				fmt.Println(1111)
				time.Sleep(1 * time.Second)
			}
			fmt.Println("Sleep After")
		}
	}
}
