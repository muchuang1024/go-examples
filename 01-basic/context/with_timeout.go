package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		// 执行耗时的函数
		time.Sleep(5 * time.Second)
		fmt.Println("不会打印")
		// 在函数执行完毕后调用 cancel() 取消超时上下文
		cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Function execution timed out")
	}

	time.Sleep(100 * time.Second)
}
