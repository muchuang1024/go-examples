package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		fmt.Println(1111)
		runtime.Goexit()
		fmt.Println(22222)
	}()
	time.Sleep(1 * time.Second)
}
