package main

import "fmt"

func F() {
	defer fmt.Println(111)
	panic("a")
	defer fmt.Println(222)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
		fmt.Println("c")
	}()
	F()
	fmt.Println("继续执行")
}
