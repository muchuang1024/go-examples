package main

import "fmt"

func main() {
	floor := 4 - (10-1)/9
	column := (10-1)%9 + 1
	fmt.Println(floor, column)
	node := []int{1}
	slice1 := [][]int{node}
	slice2 := make([][]int, len(slice1))
	slice2[0] = make([]int, len(slice1[0]))
	copy(slice2[0], slice1[0])
	fmt.Println(slice1, slice2)
	node[0] = 2
	fmt.Println(slice1, slice2)
}
