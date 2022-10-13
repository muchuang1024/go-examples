package main

import "fmt"

// go tool compile -N -l -S compile.go

func main() {
	a := 1
	defer fmt.Println("the value of a1:", a)
	a++

	defer func() {
		fmt.Println("the value of a2:", a)
	}()

	m := make(map[string]int, 0)
	m["test"] = 1
	fmt.Println(m)
	v := m["test"]
	fmt.Println(v)
	var s []int = []int{1, 2}
	fmt.Println(s)
}