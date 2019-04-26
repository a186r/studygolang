// // // // // // // 在go中，与c数组变量隐式作为指针使用不同，Go数组是值类型，赋值和函数传参都会复制整个数组数据

// // // // // // package main

// // // // // // import "fmt"

// // // // // // func main() {
// // // // // // 	arrayA := [2]int{100, 200}
// // // // // // 	var arrayB [2]int

// // // // // // 	arrayB = arrayA
// // // // // // 	fmt.Printf("arrayA : %p, %v\n", &arrayA, arrayA)
// // // // // // 	fmt.Printf("arrayB : %p, %v\n", &arrayB, arrayB)

// // // // // // 	testArray(arrayA)
// // // // // // }

// // // // // // func testArray(x [2]int) {
// // // // // // 	fmt.Printf("func Array : %p, %v\n", &x, x)
// // // // // // }

// // // // // // // 打印可以看到三个数组的内存地址都不同，这也验证了Go中数组的赋值和传参都是值复制的

// // // // // // // 这样就有一个问题，如果数组大小有100w，在64位及其上就需要花费大约800w字节，即8M内存，这样会消耗大量的内存。
// // // // // // // 于是有人想到，函数传参用数组的指针

// // // // // package main

// // // // // import "fmt"

// // // // // func main() {
// // // // // 	arrayA := [2]int{100, 200}
// // // // // 	testArrayPoint(&arrayA) // 1.传数组指针
// // // // // 	arrayB := arrayA[:]
// // // // // 	testArrayPoint(&arrayB) // 2.传切片
// // // // // 	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
// // // // // }

// // // // // func testArrayPoint(x *[]int) {
// // // // // 	fmt.Printf("func Array : %p , %v\n", x, *x)
// // // // // 	(*x)[1] += 100
// // // // // }

// // // // // 切片本身是一个只读对象，其工作机制类似数组指针的一种封装
// // // // package main

// // // // import "unsafe"

// // // // // slice的数据结构定义如下
// // // // type slice struct {
// // // // 	array unsafe.Pointer
// // // // 	len   int
// // // // 	cap   int
// // // // }

// // // // func mian() {
// // // // 	// // 如果想从slice中得到一块内存地址，可以这样做：
// // // // 	// s := make([]byte, 200)
// // // // 	// ptr := unsafe.Pointer(&s[0])

// // // // 	// 如果反过来呢，从Go的内存地址中构造一个slice
// // // // 	var ptr unsafe.Pointer
// // // // 	var s1 = struct {
// // // // 		addr uintptr
// // // // 		len  int
// // // // 		cap  int
// // // // 	}{ptr, length, length}

// // // // 	s := *(*[]byte)(unsafe.Pointer(&s1))
// // // // }

// // // // 创建slice
// // // package mian

// // // func main() {

// // // }

// // // func makeslice(et *type, len, cap int) slice {
// // // 	// 根据切片的数据类型，获取切片的最大容量
// // // 	maxElements := maxSliceCap(et.size)
// // // 	// 比较切片的长度，长度值域应该在[0,maxElements]之间
// // // 	if len < 0 || uintptr(len) > maxElements {
// // // 		panic(errorString("makeslice: len out of range"))
// // // 	}

// // // 	// 比较切片的容量，容量阈值应该在[len,maxElements]之间
// // // 	if cap < len || uintptr(cap) >maxElements{
// // // 		panic(errorString("makeslice: cap out of range"))
// // // 	}
// // // 	// 根据切片的容量申请内存
// // // 	p := mallocgc(et.size*uintptr(cap), et, true)
// // // 	// 返回申请好内存的切片的首地址
// // // 	return slice{p,len,cap}
// // // }

// // package mian

// // func main() {
// // 	// 除了make可以创建slice，字面量也可以创建slice
// // 	// 这里用字面量创建一个len=6，cap=6的切片，这个时候数组里面的每个元素的值都初始化完成了
// // 	// 需要注意的是[]里面不要写数组的容量，写了以后就是数组了，而不是切片了
// // 	slice := []int{10, 20, 30, 40, 50} //切片
// // 	slice1 := [5]int{1, 2, 3, 4, 5}    //数组
// // }

// // nil切片和空切片

// // nil切片被用在很多标准库和内置函数中，描述一个不存在的切片的时候，就需要用到nil切片。
// // 比如函数在发生异常的时候，返回的切片就是nil切片

// // 空切片，空切片也很常用，一般会来表示一个空的集合，比如数据库查询，一条结果也没查到，就可以返回一个空切片
// package mian

// func main() {
// 	// nil切片
// 	// var slice []int

// 	// 空切片
// 	// slice := make([]int,0)

// }
