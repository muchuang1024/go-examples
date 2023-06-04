package main

import (
	"fmt"
	"time"
)

// 非缓存channel只写不读
func deadlock1() {
	ch := make(chan int)
	ch <- 3
}

// 非缓存channel读在写后面
func deadlock2_1() {
	ch := make(chan int)
	ch <- 3
	num := <-ch
	fmt.Println("num=", num)
}

// 非缓存channel读在写后面
func deadlock2_2() {
	ch := make(chan int)
	ch <- 100
	go func() {
		num := <-ch
		fmt.Println("num=", num)
	}()
	time.Sleep(time.Second)
}

// 写入超过缓冲区数量（缓冲channel）
func deadlock3() {
	ch := make(chan int, 3)
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
}

// 空读导致死锁
func deadlock4() {
	ch := make(chan int)
	// ch := make(chan int, 1)
	fmt.Println(<-ch)
}

// 多个协程互相等待
func deadlock5() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println("num=", num)
				ch2 <- 100
			}
		}
	}()

	for {
		select {
		case num := <-ch2:
			fmt.Println("num=", num)
			ch1 <- 300
		}
	}
}

func main() {
	deadlock5()
}
