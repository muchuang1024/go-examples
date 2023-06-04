package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// 模拟用户请求数量
	requestCount := 10
	fmt.Println("goroutine_num", runtime.NumGoroutine())
	// 管道长度即最大并发数
	ch := make(chan bool, 3)
	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		ch <- true
		go Read(ch, i)
	}

	wg.Wait()
}

func Read(ch chan bool, i int) {
	fmt.Printf("goroutine_num: %d, go func: %d\n", runtime.NumGoroutine(), i)
	<-ch
	wg.Done()
}
