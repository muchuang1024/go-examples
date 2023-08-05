package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())

	go func() {
		// 执行耗时的函数
		time.Sleep(5 * time.Second)
		fmt.Println("不会打印")

		// 在函数执行完毕后调用 cancel() 取消上下文
		// cancel()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("Function execution timed out")
	}
}
