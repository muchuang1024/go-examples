package main

func single() {
	ch1 := make(chan<- int)
	ch2 := make(<-chan int)

}

func main() {
	single()
}
