package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	for i := 0; i < 10; i++ {
		v := pool.Get().(int)
		fmt.Println(v) // 取出来的值是put进去的，对象复用；如果是新建对象，则取出来的值为0
		pool.Put(i)
	}
}
