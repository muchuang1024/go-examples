package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	wg.Add(3)
	start := time.Now().Unix()
	go print("gorouine1", ch1, ch2)
	go print("gorouine2", ch2, ch3)
	go print("gorouine3", ch3, ch1)
	ch1 <- struct{}{}
	wg.Wait()
	end := time.Now().Unix()
	fmt.Printf("duration:%d\n", end-start)
}

func print(gorouine string, inputchan chan struct{}, outchan chan struct{}) {
	// 模拟内部操作耗时
	time.Sleep(1 * time.Second)
	for {
		select {
		case <-inputchan:
			fmt.Printf("%s\n", gorouine)
			outchan <- struct{}{}
			wg.Done()
			return
		default:
			break
		}
	}
}
