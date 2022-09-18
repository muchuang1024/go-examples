package main

import "fmt"

func main() {
	var args int64 = 1                  // int类型变量
	p := &args                          // 指针类型变量
	fmt.Printf("原始指针的内存地址是 %p\n", &p)   // 存放指针类型变量
	fmt.Printf("原始指针指向变量的内存地址 %p\n", p) // 存放int变量
	modifyPointer(p)                    // args就是实际参数
	fmt.Printf("改动后的值是: %v\n", *p)
}

func modifyPointer(p *int64) { //这里定义的args就是形式参数
	fmt.Printf("函数里接收到指针的内存地址是 %p \n", &p)
	fmt.Printf("函数里接收到指针指向变量的内存地址 %p\n", p)
	*p = 10
}
