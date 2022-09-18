package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 封装好的计数器
	var counter Counter
	var gNum = 1000
	// 启动10个goroutine
	for i := 0; i < gNum; i++ {
		go func() {
			counter.Count() // 受到锁保护的方法
		}()
	}
	for { // 一个writer
		counter.Incr() // 计数器写操作
		fmt.Println("incr")
		time.Sleep(time.Second)
	}
}

// 线程安全的计数器类型
type Counter struct {
	mu    sync.RWMutex
	count uint64
}

// 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 得到计数器的值，也需要锁保护
func (c *Counter) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
