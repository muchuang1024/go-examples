package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(1111, len(s), s[6:])
	a := make(map[int]bool)
	a = map[int]bool{
		1: true,
		2: true,
	}
	for fromAreaId, _ := range a {
		for toAreaId, _ := range a {
			if fromAreaId == toAreaId {
				continue
			}
			fmt.Println(fromAreaId, toAreaId)
		}
	}
	fmt.Println(rand.Intn(2))
	fmt.Println(rand.Intn(2))
	// a := []int{1, 2, 3, 4, 5}
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Intn(5))
	// }
	// fmt.Println(a[1:4], len(a), a[5:])
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Intn(2))
	// }
	// a := []int{1, 2, 3, 4}
	// fmt.Println(singleNumbers(a))
	// b := []int{1, 2, 3, 3}
	// fmt.Println(singleNumbers(b))
}

func singleNumbers(nums []int) []int {
	ret := 0
	for _, val := range nums {
		ret ^= val
	}
	pointer := 1
	for pointer&ret == 0 { // 因为a，b相同的位的异或结果必定是0
		pointer = pointer << 1
	}
	a, b := 0, 0
	for _, val := range nums {
		if pointer&val != 0 {
			fmt.Println(111, val)
		} else {
			fmt.Println(222, val)
		}
	}
	return []int{a, b}
}
