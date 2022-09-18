package main

import (
	"fmt"
	"time"
)

func F() {
	defer fmt.Println(111)
	panic("a")
	defer fmt.Println(222)
}

func main() {
	defer func() {
		// 无法捕获其它goroutine的panic
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("c")
	}()
	go F() // 外层无法捕获
	time.Sleep(time.Second)
	fmt.Println("继续执行")
}

// 输出 捕获异常: a -> b -> 继续执行 -> c
