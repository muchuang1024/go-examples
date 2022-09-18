package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

// nil channel
func block1() {
	var ch chan int
	for i := 0; i < 10; i++ {
		go func() {
			<-ch
		}()
	}
}

// 发送不接收
func block2() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			ch <- 1
		}()
	}
}

// 接收不发送
func block3() {
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			<-ch
		}()
	}
}

func requestWithNoClose() {
	_, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("error occurred while fetching page, error: %s", err.Error())
	}
}

func requestWithClose() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("error occurred while fetching page, error: %s", err.Error())
		return
	}
	defer resp.Body.Close()
}

func block4() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			requestWithClose()
			// do something...
		}()
	}
}

func block5() {
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
		}()
	}
}

func block6() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(2)
			wg.Done()
			wg.Wait()
		}()
	}
}

var wg = sync.WaitGroup{}

func main() {
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	fmt.Println("before goroutines: ", runtime.NumGoroutine())
	block4()
	wg.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("after goroutines: ", runtime.NumGoroutine())
	select {}
}
