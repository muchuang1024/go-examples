package main

import "fmt"

func main() {
	var i int64 = 1
	fmt.Printf("原始int内存地址是 %p\n", &i)
	modifyInt(i) // args就是实际参数
	fmt.Printf("改动后的值是: %v\n", i)
}

func modifyInt(i int64) { //这里定义的args就是形式参数
	fmt.Printf("函数里接收到int的内存地址是：%p\n", &i)
	i = 10
}
