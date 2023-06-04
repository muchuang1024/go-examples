package main

import (
	"fmt"
	"time"
)

// 3个函数分别打印cat、dog、fish，要求每个函数都要起一个goroutine，按照cat、dog、fish顺序打印在屏幕上100次(各打印100次)。
func main() {
	for i := 0; i < 10; i++ {
		go cat()
		go dog()
		go fish()
	}
	time.Sleep(10 * time.Second)
}

func cat() {
	fmt.Println("cat")
}

func dog() {
	fmt.Println("dog")
}

func fish() {
	fmt.Println("fish")
}
