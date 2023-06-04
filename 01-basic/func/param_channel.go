package main

import (
	"fmt"
	"time"
)

func main() {
	p := make(chan bool)
	fmt.Printf("原始chan的内存地址是：%p\n", &p)
	go func(p chan bool) {
		fmt.Printf("函数里接收到chan的内存地址是：%p\n", &p)
		//模拟耗时
		time.Sleep(2 * time.Second)
		p <- true
	}(p)

	select {
	case l := <-p:
		fmt.Printf("接收到的值是: %v\n", l)
	}
}
