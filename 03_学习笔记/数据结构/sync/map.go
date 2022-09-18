package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("1", 1)
	scene.Store("2", 2)
	scene.Store("3", 3)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("1"))
	// 根据键删除对应的键值对
	scene.Delete("1")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}
