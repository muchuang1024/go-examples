package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func printGCStats() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			fmt.Printf("gc %d last@%v, PauseTotal %v\n", s.NumGC, s.LastGC, s.PauseTotal)
		}
	}
}
func main() {
	go printGCStats()
	for n := 1; n < 100000; n++ {
		_ = make([]byte, 1<<20)
	}
}
