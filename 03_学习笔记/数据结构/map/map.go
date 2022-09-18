package main

import (
	"log"
	"runtime"
)

var intMap map[int]map[int]int

var cnt = 8192

func main() {
	printMemStats()
	initMap()
	runtime.GC()
	printMemStats()
	log.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
		// intMap[i] = nil
	}
	log.Println(len(intMap))
	runtime.GC()
	printMemStats()
	intMap = nil
	runtime.GC()
	printMemStats()
}
func initMap() {
	intMap = make(map[int]map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		intMap[i] = map[int]int{}
	}
}
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}
