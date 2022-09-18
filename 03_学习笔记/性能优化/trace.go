package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

// Go 的两大杀器 pprof + trace 组合

// go run trace.go
// go tool trace trace.out

// go build trace.go
// GODEBUG=schedtrace=1000 ./trace

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello World")
	}
}
