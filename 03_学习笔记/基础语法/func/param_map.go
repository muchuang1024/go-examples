package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["age"] = 8

	fmt.Printf("原始map的内存地址是：%p\n", &m)
	modifyMap(m)
	fmt.Printf("改动后的值是: %v\n", m)
}

func modifyMap(m map[string]int) {
	fmt.Printf("函数里接收到map的内存地址是：%p\n", &m)
	m["age"] = 9
}
