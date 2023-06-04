package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var status int64

func main() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	go broadcast(c)
	time.Sleep(1 * time.Second)
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
}

func listen(c *sync.Cond) {
	c.L.Lock()
	// for atomic.LoadInt64(&status) != 1 {
	c.Wait()
	// Wait 内部会先调用 c.L.Unlock()，来先释放锁，如果调用方不先加锁的话，会报错
	// }
	fmt.Println("listen")
	c.L.Unlock()
}
