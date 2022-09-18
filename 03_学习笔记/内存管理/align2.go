package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type T1 struct {
	// i16  int16 // 2 byte
	bool bool // 1 byte
}

type T2 struct {
	i8  int8  // 1 byte
	i64 int64 // 8 byte
	i32 int32 // 4 byte
}

type T3 struct {
	i8  int8  // 1 byte
	i32 int32 // 4 byte
	i64 int64 // 8 byte
}

func main() {
	fmt.Println(runtime.GOARCH) // amd64

	t1 := T1{}
	fmt.Println(unsafe.Sizeof(t1)) // 4 bytes

	t2 := T2{}
	fmt.Println(unsafe.Sizeof(t2)) // 24 bytes

	t3 := T3{}
	fmt.Println(unsafe.Sizeof(t3)) // 16 bytes
}
