package main

import "fmt"

// go tool compile -S defer.go | grep CALL
func main() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	defer fmt.Println("defer4")
	fmt.Println("11111")
}
