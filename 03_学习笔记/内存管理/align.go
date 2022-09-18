package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("bool alignof is %d\n", unsafe.Alignof(bool(true)))
	fmt.Printf("string alignof is %d\n", unsafe.Alignof(string("a")))
	fmt.Printf("int alignof is %d\n", unsafe.Alignof(int8(0)))
	fmt.Printf("float alignof is %d\n", unsafe.Alignof(float64(0)))
	fmt.Printf("int32 alignof is %d\n", unsafe.Alignof(int32(0)))
	fmt.Printf("float32 alignof is %d\n", unsafe.Alignof(float32(0)))
}
