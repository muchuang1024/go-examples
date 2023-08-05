package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	// a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		a[k] = 100 + v
		// 0: 100, 200, 3
		// 1: 101, 200, 3
		// 2: 101, 300, 3
		// 3: 101, 300, 103
	}
	fmt.Print(a)
}
