package main

import "fmt"

type Person struct {
	age int
}

// 如果实现了接收者是指针类型的方法，会隐含地也实现了接收者是值类型的IncrAge1方法。
// 会修改age的值
func (p *Person) IncrAge1() {
	p.age += 1
}

// 如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的IncrAge2方法。
// 不会修改age的值
func (p Person) IncrAge2() {
	p.age += 1
}

// 如果实现了接收者是值类型的方法，会隐含地也实现了接收者是指针类型的GetAge方法。
func (p Person) GetAge() int {
	return p.age
}

func main() {
	// p1 是值类型
	p := Person{age: 10}

	// 值类型 调用接收者是指针类型的方法
	p.IncrAge1()
	fmt.Println(p.GetAge())
	// 值类型 调用接收者是值类型的方法
	p.IncrAge2()
	fmt.Println(p.GetAge())

	// ----------------------

	// p2 是指针类型
	p2 := &Person{age: 20}

	// 指针类型 调用接收者是指针类型的方法
	p2.IncrAge1()
	fmt.Println(p2.GetAge())
	// 指针类型 调用接收者是值类型的方法
	p2.IncrAge2()
	fmt.Println(p2.GetAge())
}
