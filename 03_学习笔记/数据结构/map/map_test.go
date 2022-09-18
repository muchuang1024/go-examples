package mymap

import (
	"sort"
	"sync"
	"testing"
	"unsafe"
)

/**
* 初始化
 */
func TestMapInit(t *testing.T) {
	t.Log(5 >> 1)
	// 初始化方式1：直接声明
	// var m1 map[string]int
	// m1["a"] = 1
	// t.Log(m1, unsafe.Sizeof(m1))
	// panic: assignment to entry in nil map
	// 向 map 写入要非常小心，因为向未初始化的 map（值为 nil）写入会引发 panic，所以向 map 写入时需先进行判空操作

	// 初始化方式2：使用字面量
	m2 := map[string]int{}
	m2["a"] = 2
	t.Log(m2, unsafe.Sizeof(m2))
	// map[a:2] 8

	// 初始化方式3：使用make创建
	m3 := make(map[string]int)
	m3["a"] = 3
	t.Log(m3, unsafe.Sizeof(m3))
	// map[a:3] 8
}

/**
* 增加
 */
func TestMapAdd(t *testing.T) {
	m := map[string]int{}
	m["a"] = 1
	m["b"] = 2
	t.Log(m)
	// map[a:1 b:2]
}

func TestMapDelete(t *testing.T) {
	m := map[string]int{"a": 1}
	delete(m, "a")
	delete(m, "b")
	t.Log(m)
	// map[]
}

func TestMapUpdate(t *testing.T) {
	m := map[string]int{"a": 1}
	m["a"] = 2
	t.Log(m)
	// map[a:2]
}

func TestMapFind(t *testing.T) {
	m := map[string]int{"a": 1}
	t.Log(m["a"])
	// 1
	t.Log(m["b"])
	// 0，key不存在时，为值类型默认值
}

func TestMapRange(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	t.Log("first range:")
	for i, v := range m {
		t.Logf("m[%v]=%v ", i, v)
	}
	t.Log("\nsecond range:")
	for i, v := range m {
		t.Logf("m[%v]=%v ", i, v)
	}

	// 实现有序遍历
	var sl []int
	// 把 key 单独取出放到切片
	for k := range m {
		sl = append(sl, k)
	}
	// 排序切片
	sort.Ints(sl)
	// 以切片中的 key 顺序遍历 map 就是有序的了
	for _, k := range sl {
		t.Log(k, m[k])
	}
}

func TestMapShareMemory(t *testing.T) {
	m1 := map[string]int{}
	m2 := m1
	m1["a"] = 1
	t.Log(m1, len(m1))
	// map[a:1] 1
	t.Log(m2, len(m2))
	// map[a:1]
}

func TestSliceFn(t *testing.T) {
	m := map[string]int{}
	t.Log(m, len(m))
	// map[a:1]
	mapAppend(m, "b", 2)
	t.Log(m, len(m))
	// map[a:1 b:2] 2
}

func mapAppend(m map[string]int, key string, val int) {
	m[key] = val
}

// 使用sync.Map并发安全，但会有一些性能损失，性能低于Mutex
func BenchmarkMapConcurrencySafeBySyncMap(b *testing.B) {
	var m sync.Map
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}
	wg.Wait()
	b.Log(b.N)
}

func BenchmarkMapConcurrencySafeByMutex(b *testing.B) {
	var lock sync.Mutex //互斥锁
	// 不加锁，会报错：fatal error: concurrent map writes
	m := make(map[int]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			m[i] = i
		}(i)
	}
	wg.Wait()
	b.Log(len(m), b.N)
}
