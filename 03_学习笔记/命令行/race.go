package main

import "fmt"

// go run -race race.go

func main() {
	i := 0

	go func() {
		i++ // write i
	}()

	fmt.Println(i) // read i
}
