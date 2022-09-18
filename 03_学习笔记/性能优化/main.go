package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

// https://darjun.github.io/2021/06/09/youdontknowgo/pprof/

// go tool pprof -http=:1248 http://127.0.0.1:6060/debug/pprof/goroutine

func main() {

	fmt.Println(68 << 1)

	for i := 0; i < 100; i++ {
		go func() {
			select {}
		}()
	}

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	select {}
}
