// 关键字
// 包声明
// 注释
// 引入包
// 函数
// 类型
// 变量和常量
// 语句
// 运算符

// 包声明
package main

// 引入包（标准库）
import (
	"fmt"
	"time"
)

func b() {
	go func() {
		for {
			fmt.Println(111)
		}
	}()
	fmt.Println(222)
}

// 函数
func main() {
	go func() {
		b()
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(1 ^ 0)
	fmt.Println(0 ^ 1)
	fmt.Println(1 ^ 1)
	fmt.Println(0 ^ 0)
	// 变量类型声明
	var name string
	// 语句（赋值语句、条件语句、循环语句、跳转语句）
	name = "gopher"
	// 运算符（算术运算、位运算、逻辑运算）
	fmt.Println("hello world, " + name)
}
