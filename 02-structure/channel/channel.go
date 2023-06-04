package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan struct{}, 1)
	for i := 0; i < 10; i++ {
		go func() {
			c <- struct{}{}
			time.Sleep(1 * time.Second)
			fmt.Println("通过ch访问临界区")
			<-c
		}()
	}
	for {
	}

	ch := make(chan string)
	go sendTask(ch)
	go receiveTask(ch)
	time.Sleep(1 * time.Second)
}

//G1
func sendTask(ch chan string) {
	taskList := []string{"this", "is", "a", "demo"}
	// 阻塞方式
	for _, task := range taskList {
		ch <- task //发送任务到channel
	}
	// 非阻塞方式
}

//G2
func receiveTask(ch chan string) {
	for {
		task := <-ch                  //接收任务
		fmt.Println("received", task) //处理任务
	}
}
