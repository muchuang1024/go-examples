package main

import (
	"fmt"
)

func notInitChannel() {
	var ch chan int
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
}

func closedChannel() {

	ch := make(chan int, 3)
	go func() {
		ch <- 1
		ch <- 2
	}()
	fmt.Println(<-ch)
	close(ch)
	v, ok := <-ch
	fmt.Println(v, ok)
	v2, ok := <-ch
	fmt.Println(v2, ok)
	v3, ok := <-ch
	fmt.Println(v3, ok)
}

func main() {
	closedChannel()
}
