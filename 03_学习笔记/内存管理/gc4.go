package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMemStats() {
	t := time.NewTicker(time.Second)
	s := runtime.MemStats{}
	for {
		select {
		case <-t.C:
			runtime.ReadMemStats(&s)
			fmt.Printf("gc %d last@%v, heap_object_num: %v, heap_alloc: %vMB, next_heap_size: %vMB\n",
				s.NumGC, time.Unix(int64(time.Duration(s.LastGC).Seconds()), 0), s.HeapObjects, s.HeapAlloc/(1<<20), s.NextGC/(1<<20))
		}
	}
}
func main() {
	go printMemStats()
	for n := 1; n < 100000; n++ {
		_ = make([]byte, 1<<20)
	}
}
