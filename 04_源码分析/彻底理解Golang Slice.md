![](https://cdn.jsdelivr.net/gh/caijinlin/imgcdn/%E6%B7%B1%E5%85%A5%E7%90%86%E8%A7%A3Slice.png)

>  看完这篇文章，下面这些高频面试题你都会答了吧

1. Go slice的底层实现原理
2. Go array和slice的区别
3. Go slice深拷贝和浅拷贝
4. Go slice扩容机制是怎样的？
5. 为什么Go slice是非线程安全的？

## 实现原理

slice是无固定长度的数组，底层结构是一个结构体，包含如下3个属性

一个 `slice` 在 golang 中占用 24 个 bytes

```
type slice struct {
	array unsafe.Pointer 
	len   int 
	cap   int 
}
```

array : 包含了一个指向一个数组的指针，数据实际上存储在这个指针指向的数组上，占用 8 bytes

len:  当前 slice 使用到的长度，占用8 bytes

cap :  当前 slice 的容量，同时也是底层数组 array 的长度， 8 bytes



![](https://cdn.jsdelivr.net/gh/caijinlin/imgcdn/1605449197795-9ca02de7-b129-4f7c-a0c3-930798b881c0.svg)



slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像 array一样，只是长度可变。**golang中通过语法糖，使得我们可以像声明array一样，自动创建slice结构体**

*根据*索引位置取切片*slice* 元素值时，默认取值范围是（0～*len*(*slice*)-1），一般输出slice时，通常是指 slice[0:len(slice)-1]，根据下标就可以输出所指向底层数组中的值

## 主要特性

### 引用类型

golang 有三个常用的高级类型*slice*、map、channel, 它们都是*引用类型*，当引用类型作为函数参数时，可能会修改原内容数据。

```
func sliceModify(s []int) {
	s[0] = 100
}

func sliceAppend(s []int) []int {
	s = append(s, 100)
	return s
}

func sliceAppendPtr(s *[]int) {
	*s = append(*s, 100)
	return
}

// 注意：Go语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
// 拷贝的内容是非引用类型（int、string、struct等这些），在函数中就无法修改原内容数据；
// 拷贝的内容是引用类型（interface、指针、map、slice、chan等这些），这样就可以修改原内容数据。
func TestSliceFn(t *testing.T) {
	// 参数为引用类型slice：外层slice的len/cap不会改变，指向的底层数组会改变
	s := []int{1, 1, 1}
	newS := sliceAppend(s)
	// 函数内发生了扩容
	t.Log(s, len(s), cap(s))
	// [1 1 1] 3 3
	t.Log(newS, len(newS), cap(newS)) 
	// [1 1 1 100] 4 6

	s2 := make([]int, 0, 5)
	newS = sliceAppend(s2)
	// 函数内未发生扩容
	t.Log(s2, s2[0:5], len(s2), cap(s2)) 
	// [] [100 0 0 0 0] 0 5
	t.Log(newS, newS[0:5], len(newS), cap(newS))
	// [100] [100 0 0 0 0] 1 5

	// 参数为引用类型slice的指针：外层slice的len/cap会改变，指向的底层数组会改变
	sliceAppendPtr(&s)
	t.Log(s, len(s), cap(s)) 
  // [1 1 1 100] 4 6
	sliceModify(s)
	t.Log(s, len(s), cap(s)) 
  // [100 1 1 100] 4 6
}
```

公众号后台caspar回复【代码】获取本文所有示例代码

### 切片状态

切片有3种特殊的状态，分为「零切片」、「空切片」和「nil 切片」

```
func TestSliceEmptyOrNil(t *testing.T) {
	var slice1 []int           
  // slice1 is nil slice
	slice2 := make([]int, 0)    
	// slcie2 is empty slice
	var slice3 = make([]int, 2) 
	// slice3 is zero slice
	if slice1 == nil {
		t.Log("slice1 is nil.") 
		// 会输出这行
	}
	if slice2 == nil {
		t.Log("slice2 is nil.") 
		// 不会输出这行
	}
	t.Log(slice3) // [0 0]
}
```

### 非线程安全

slice不支持并发读写，所以并不是线程安全的，使用多个 goroutine 对类型为 slice 的变量进行操作，每次输出的值大概率都不会一样，与预期值不一致;  slice在并发执行中不会报错，但是数据会丢失

```
/**
* 切片非并发安全
* 多次执行，每次得到的结果都不一样
* 可以考虑使用 channel 本身的特性 (阻塞) 来实现安全的并发读写
 */
func TestSliceConcurrencySafe(t *testing.T) {
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			a = append(a, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log(len(a)) 
	// not equal 10000
}

```

如果想实现slice线程安全，有两种方式：

方式一：通过加锁实现slice线程安全，适合对性能要求不高的场景。

```
func TestSliceConcurrencySafeByMutex(t *testing.T) {
	var lock sync.Mutex //互斥锁
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			a = append(a, i)
		}(i)
	}
	wg.Wait()
	t.Log(len(a)) 
	// equal 10000
}
```

方式二：通过channel实现slice线程安全，适合对性能要求高的场景。

```
func TestSliceConcurrencySafeByChanel(t *testing.T) {
	buffer := make(chan int)
	a := make([]int, 0)
	// 消费者
	go func() {
		for v := range buffer {
			a = append(a, v)
		}
	}()
	// 生产者
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			buffer <- i
		}(i)
	}
	wg.Wait()
	t.Log(len(a)) 
	// equal 10000
}
```

### 共享存储空间

多个切片如果共享同一个底层数组，这种情况下，对其中一个切片或者底层数组的更改，会影响到其他切片

```
/**
* 切片共享存储空间
 */
func TestSliceShareMemory(t *testing.T) {
	slice1 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"}
	Q2 := slice1[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) 
	// [4 5 6] 3 9
	Q3 := slice1[5:8]
	t.Log(Q3, len(Q3), cap(Q3)) 
	// [6 7 8] 3 7
	Q3[0] = "Unkown"
	t.Log(Q2, Q3) 
	// [4 5 Unkown] [Unkown 7 8]

	a := []int{1, 2, 3, 4, 5}
	shadow := a[1:3]
	t.Log(shadow, a)             
	// [2 3] [1 2 3 4 5]
	shadow = append(shadow, 100) 
	// 会修改指向数组的所有切片
	t.Log(shadow, a)            
  // [2 3 100] [1 2 3 100 5]
}
```

## 常用操作

### 创建

slice 的创建有4种方式，如下：

```
func TestSliceInit(t *testing.T) {
	// 初始化方式1：直接声明
	var slice1 []int
	t.Log(len(slice1), cap(slice1)) 
	// 0, 0
	slice1 = append(slice1, 1)
	t.Log(len(slice1), cap(slice1)) 
	// 1, 1, 24

	// 初始化方式2：使用字面量
	slice2 := []int{1, 2, 3, 4}
	t.Log(len(slice2), cap(slice2)) 
	// 4, 4, 24

	// 初始化方式3：使用make创建slice
	slice3 := make([]int, 3, 5)           
  // make([]T, len, cap) cap不传则和len一样
	t.Log(len(slice3), cap(slice3))       
  // 3, 5
	t.Log(slice3[0], slice3[1], slice3[2]) 
	// 0, 0, 0
	// t.Log(slice3[3], slice3[4]) 
	// panic: runtime error: index out of range [3] with length 3
	slice3 = append(slice3, 1)
	t.Log(len(slice3), cap(slice3)) 
	// 4, 5, 24

	// 初始化方式4: 从切片或数组“截取”
	arr := [100]int{}
	for i := range arr {
		arr[i] = i
	}
	slcie4 := arr[1:3]
	slice5 := make([]int, len(slcie4))
	copy(slice5, slcie4)
	t.Log(len(slcie4), cap(slcie4), unsafe.Sizeof(slcie4)) 
	// 2，99，24
	t.Log(len(slice5), cap(slice5), unsafe.Sizeof(slice5)) 
	// 2，2，24
}
```

### 增加

```
func TestSliceGrowing(t *testing.T) {
	slice1 := []int{}
	for i := 0; i < 10; i++ {
		slice1 = append(slice1, i)
		t.Log(len(slice1), cap(slice1))
	}
	// 1 1
	// 2 2
	// 3 4
	// 4 4
	// 5 8
	// 6 8
	// 7 8
	// 8 8
	// 9 16
	// 10 16
}
```

### 删除

```
func TestSliceDelete(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	var x int
	// 删除最后一个元素
	x, slice1 = slice1[len(slice1)-1], slice1[:len(slice1)-1] 
	t.Log(x, slice1, len(slice1), cap(slice1)) 
	// 5 [1 2 3 4] 4 5

	// 删除第2个元素	
	slice1 = append(slice1[:2], slice1[3:]...) 
	t.Log(slice1, len(slice1), cap(slice1))    
	// [1 2 4] 3 5
}
```

### 查找

```
v := s[i] // 下标访问
```

### 修改

```
s[i] = 5 // 下标修改
```

### 截取

```
/**
* 切片截取
 */
func TestSliceSubstr(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[:]
	// 截取 slice[left:right:max]
	// left：省略默认0
	// right：省略默认len(slice1)
	// max: 省略默认len(slice1)
	// len = right-left+1
	// cap = max-left
	t.Log(slice2, len(slice2), cap(slice2)) 
	// 1 2 3 4 5] 5 5
	slice3 := slice1[1:]
	t.Log(slice3, len(slice3), cap(slice3)) 
	// [2 3 4 5] 4 4
	slice4 := slice1[:2]
	t.Log(slice4, len(slice4), cap(slice4)) 
	// [1 2] 2 5
	slice5 := slice1[1:2]
	t.Log(slice5, len(slice5), cap(slice5)) 
	// [2] 1 4
	slice6 := slice1[:2:5]
	t.Log(slice6, len(slice6), cap(slice6)) 
	// [1 2] 2 5
	slice7 := slice1[1:2:2]
	t.Log(slice7, len(slice7), cap(slice7)) 
	// [2] 1 1
}
```

### 遍历

切片有3种遍历方式

```
/**
* 切片遍历
 */
func TestSliceTravel(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	for i := 0; i < len(slice1); i++ {
		t.Log(slice1[i])
	}
	for idx, e := range slice1 {
		t.Log(idx, e)
	}
	for _, e := range slice1 {
		t.Log(e)
	}
}
```

### 反转

```
func TestSliceReverse(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	t.Log(a, len(a), cap(a)) 
	// [5 4 3 2 1] 5 5
}
```

### 拷贝

开发中会经常的把一个变量复制给另一个变量，那么这个过程，可能是深浅拷贝，那么今天帮大家区分一下这两个拷贝的区别和具体的区别

#### 深拷贝

拷贝的是数据本身，创造一个样的新对象，新创建的对象与原对象不共享内存，新创建的对象在内存中开辟一个新的内存地址，新对象值修改时不会影响原对象值。既然内存地址不同，释放内存地址时，可分别释放

值类型的数据，默认赋值操作都是深拷贝，Array、Int、String、Struct、Float，Bool。引用类型的数据如果想实现深拷贝，需要通过辅助函数完成

比如golang深拷贝copy 方法会把源切片值(即 from Slice )中的元素复制到目标切片(即 to Slice )中，并返回被复制的元素个数，copy 的两个类型必须一致。copy 方法最终的**复制结果取决于较短的那个切片**，当较短的切片复制完成，整个复制过程就全部完成了

```
/**
* 深拷贝
 */
func TestSliceDeepCopy(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5, 5)
	// 深拷贝
	copy(slice2, slice1)                   
	t.Log(slice1, len(slice1), cap(slice1)) 
	// [1 2 3 4 5] 5 5
	t.Log(slice2, len(slice2), cap(slice2)) 
	// [1 2 3 4 5] 5 5
	slice1[1] = 100                        
	t.Log(slice1, len(slice1), cap(slice1)) 
	// [1 100 3 4 5] 5 5
	t.Log(slice2, len(slice2), cap(slice2)) 
	// [1 2 3 4 5] 5 5
}
```

#### 浅拷贝

拷贝的是数据地址，只复制指向的对象的指针，此时新对象和老对象指向的内存地址是一样的，新对象值修改时老对象也会变化。释放内存地址时，同时释放内存地址。

引用类型的数据，默认全部都是浅拷贝，Slice、Map等

目的切片和源切片指向同一个底层数组，任何一个数组元素改变，都会同时影响两个数组。

```
/**
* 浅拷贝
 */
func TestSliceShadowCopy(t *testing.T) {
	slice1 := []int{1, 2, 3, 4, 5}
	// 浅拷贝（注意 := 对于引用类型是浅拷贝，对于值类型是深拷贝）
	slice2 := slice1     
	t.Logf("%p", slice1) // 0xc00001c120
	t.Logf("%p", slice2) // 0xc00001c120
	// 同时改变两个数组，这时就是浅拷贝，未扩容时，修改 slice1 的元素之后，slice2 的元素也会跟着修改
	slice1[0] = 10
	t.Log(slice1, len(slice1), cap(slice1)) 
	// [10 2 3 4 5] 5 5
	t.Log(slice2, len(slice2), cap(slice2)) 
	// [10 2 3 4 5] 5 5
	// 注意下：扩容后，slice1和slice2不再指向同一个数组，修改 slice1 的元素之后，slice2 的元素不会被修改了
	slice1 = append(slice1, 5, 6, 7, 8)
	slice1[0] = 11   
  // 这里可以发现，slice1[0] 被修改为了 11, slice1[0] 还是10
	t.Log(slice1, len(slice1), cap(slice1)) 
	// [11 2 3 4 5 5 6 7 8] 9 10
	t.Log(slice2, len(slice2), cap(slice2))
  // [10 2 3 4 5] 5 5
}
```

**在复制 slice 的时候，slice 中数组的指针也被复制了，在触发扩容逻辑之前，两个 slice 指向的是相同的数组，触发扩容逻辑之后指向的就是不同的数组了**

## 扩容

扩容会发生在slice append的时候，当slice的cap不足以容纳新元素，就会进行扩容

源码：https://github.com/golang/go/blob/master/src/runtime/slice.go

```
func growslice(et *_type, old slice, cap int) slice {
	  // 省略一些判断...
    newcap := old.cap
    doublecap := newcap + newcap
    if cap > doublecap {
        newcap = cap
    } else {
        if old.len < 1024 {
            newcap = doublecap
        } else {
            // Check 0 < newcap to detect overflow
            // and prevent an infinite loop.
            for 0 < newcap && newcap < cap {
                newcap += newcap / 4
            }
            // Set newcap to the requested cap when
            // the newcap calculation overflowed.
            if newcap <= 0 {
                newcap = cap
            }
        }
    }
    // 省略一些后续...
}
```

- 如果新申请容量比两倍原有容量大，那么扩容后容量大小 等于 新申请容量
- 如果原有 slice 长度小于 1024， 那么每次就扩容为原来的 2 倍
- 如果原 slice 大于等于 1024， 那么每次扩容就扩为原来的 1.25 倍

# 内存泄露

由于slice的底层是数组，很可能数组很大，但slice所取的元素数量却很小，这就导致数组占用的绝大多数空间是被浪费的

Case1:

比如下面的代码，如果传入的`slice b`是很大的，然后引用很小部分给全局量`a`，那么`b`未被引用的部分（下标1之后的数据）就不会被释放，造成了所谓的内存泄漏。

```
var a []int

func test(b []int) {
	a = b[:1] // 和b共用一个底层数组
	return
}
```

那么只要全局量`a`在，`b`就不会被回收。

如何避免？

在这样的场景下注意：如果我们只用到一个slice的一小部分，那么底层的整个数组也将继续保存在内存当中。当这个底层数组很大，或者这样的场景很多时，可能会造成内存急剧增加，造成崩溃。 

所以在这样的场景下，我们可以将需要的分片复制到一个新的slice中去，减少内存的占用

```
var a []int

func test(b []int) {
	a = make([]int, 1)
	copy(a, b[:0])
	return
}
```

Case2:

比如下面的代码，返回的slice是很小一部分，这样该函数退出后，原来那个体积较大的底层数组也无法被回收

```
func test2() []int{
	s = make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		s = append(s, p)
	}
	s2 := s[100:102]
	return s2
}
```

如何避免？

将需要的分片复制到一个新的slice中去，减少内存的占用

```
func test2() []int{
	s = make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
	  // 一些计算...
		s = append(s, p)
	}
	s2 := make([]int, 2)
	copy(s2, s[100:102])
	return s2
}
```

# 切片与数组对比

数组是一个固定长度的，初始化时候必须要指定长度，不指定长度的话就是切片了

数组是值类型，将一个数组赋值给另一个数组时，传递的是一份深拷贝，赋值和函数传参操作都会复制整个数组数据，会占用额外的内存；切片是引用类型，将一个切片赋值给另一个切片时，传递的是一份浅拷贝，赋值和函数传参操作只会复制len和cap，但底层共用同一个数组，不会占用额外的内存。

```
//a是一个数组，注意数组是一个固定长度的，初始化时候必须要指定长度，不指定长度的话就是切片了
a := [3]int{1, 2, 3}
//b是数组，是a的一份深拷贝
b := a
//c是切片，是引用类型，底层数组是a
c := a[:]
for i := 0; i < len(a); i++ {
 a[i] = a[i] + 1
}
//改变a的值后，b是a的拷贝，b不变，c是引用，c的值改变
fmt.Println(a) 
//[2,3,4]
fmt.Println(b) 
//[1 2 3]
fmt.Println(c) 
//[2,3,4]
```

```
//a是一个切片，不指定长度的话就是切片了
a := []int{1, 2, 3}
//b是切片，是a的一份拷贝
b := a
//c是切片，是引用类型
c := a[:]
for i := 0; i < len(a); i++ {
 a[i] = a[i] + 1
}
//改变a的值后，b是a的浅拷贝，b的值改派，c是引用，c的值改变
fmt.Println(a) 
//[2,3,4]
fmt.Println(b) 
//[2,3,4]
fmt.Println(c) 
//[2,3,4]
```

# 总结

- 创建切片时可根据实际需要预分配容量，尽量避免追加过程中进行扩容操作，有利于提升性能
- 使用 append() 向切片追加元素时有可能触发扩容，扩容后将会生成新的切片
- 使用 len()、cap()计算切片长度、容量时，时间复杂度均为 O(1)，不需要遍历切片
- 切片是非线程安全的，如果要实现线程安全，可以加锁或者使用Channel
- 大数组作为函数参数时，会复制整个数组数据，消耗过多内存，建议使用切片或者指针
- 切片作为函数参数时，可以改变切片指向的数组，不能改变切片本身len和cap；想要改变切片本身，可以将改变后的切片返回 或者 将**切片指针**作为函数参数。
- 如果只用到大slice的一小部分，建议将需要的分片复制到一个新的slice中去，减少内存的占用