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
	ch := make(chan bool)
	for i := 0; i < 3; i++ {
		go Read(ch, i)
	}

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		ch <- true
	}

	wg.Wait()
}

func Read(ch chan bool, i int) {
	for _ = range ch {
		fmt.Printf("goroutine_num: %d, go func: %d\n", runtime.NumGoroutine(), i)
		wg.Done()
	}
}
