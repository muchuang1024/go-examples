package main

import "fmt"

func main() {
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	defer fmt.Println("defer4")
	fmt.Println("11111")
}

// 11111
// defer4
// defer3
// defer2
// defer1
