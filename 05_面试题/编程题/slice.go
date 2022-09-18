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
	}
	fmt.Print(a)
}
