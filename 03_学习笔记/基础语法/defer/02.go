package main

import "fmt"

// panic后的defer函数不会被执行
func main() {
	defer fmt.Println("panic before")
	panic("发生panic")
	defer func() {
		fmt.Println("panic after")
	}()
}

// panic before
// panic: 发生panic
