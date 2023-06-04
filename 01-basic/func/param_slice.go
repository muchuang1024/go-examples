package main

import "fmt"

func main() {
	var s = []int64{1, 2, 3}
	// &操作符打印出的地址是无效的，是fmt函数作了特殊处理
	fmt.Printf("直接对原始切片取地址%v \n", &s)
	// 打印slice的内存地址是可以直接通过%p打印的,不用使用&取地址符转换
	fmt.Printf("原始切片的内存地址： %p \n", s)
	fmt.Printf("原始切片第一个元素的内存地址： %p \n", &s[0])
	modifySlice(s)
	fmt.Printf("改动后的值是: %v\n", s)
}

func modifySlice(s []int64) {
	// &操作符打印出的地址是无效的，是fmt函数作了特殊处理
	fmt.Printf("直接对函数里接收到切片取地址%v\n", &s)
	// 打印slice的内存地址是可以直接通过%p打印的,不用使用&取地址符转换
	fmt.Printf("函数里接收到切片的内存地址是 %p \n", s)
	fmt.Printf("函数里接收到切片第一个元素的内存地址： %p \n", &s[0])
	s[0] = 10
}
