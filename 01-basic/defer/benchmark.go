package main

import (
	"sync"
	"testing"
)

type channel chan int

func NoDefer() {
	ch1 := make(channel, 1)
	close(ch1)
}

func Defer() {
	ch2 := make(channel, 1)
	defer close(ch2)
}

func NoDeferLock() {
	l := sync.Mutex{}
	l.Lock()
	l.Unlock()
}

func DeferLock() {
	l := sync.Mutex{}
	l.Lock()
	defer l.Unlock()
}

func BenchmarkNoDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoDefer()
	}
}

func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Defer()
	}
}

func BenchmarkNoDeferLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NoDeferLock()
	}
}

func BenchmarkDeferLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DeferLock()
	}
}
