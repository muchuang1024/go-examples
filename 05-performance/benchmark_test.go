package learn

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"
)

// go test -bench=. -benchtime=3s -run=none
func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d,%d", 2211103, 2211104)
	}
}

func BenchmarkBite(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		riderByte := make([]byte, 64)
		binary.LittleEndian.PutUint32(riderByte[0:32], 2211103)
		binary.LittleEndian.PutUint32(riderByte[32:], 2211104)
	}
}

func BenchmarkFormat(b *testing.B) {
	num := int64(10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(num, 10)
	}
}

func BenchmarkItoa(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strconv.Itoa(num)
	}
}
