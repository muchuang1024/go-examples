package main

import "fmt"

func escape6() {
	number := 10
	s := make([]int, number)
	for i := 0; i < len(s); i++ {
		s[i] = i
	}
}

func escape5() *int {
	var a int = 1
	return &a
}

func escape4() func() int {
	var i int = 1
	return func() int {
		i++
		return i
	}
}

func escape3() {
	fmt.Println(1111)
}

func escape2() {
	s := make([]int, 0, 8192)
	for index, _ := range s {
		s[index] = index
	}
}

func escape1(x, y int) *int {
	res := 0
	res = x + y
	return &res
}

func main() {
	escape6()
}
