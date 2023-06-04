package main

import "fmt"

func F() {
	defer func() {
		fmt.Println("b")
	}()
	panic("a")
}

// 子函数抛出的panic没有recover时，上层函数时，程序直接异常终止
func main() {
	defer func() {
		fmt.Println("c")
	}()
	F() // 等同于函数替换
	fmt.Println("继续执行")
}

// b
// c
// panic: a
