// // package main

// // import "fmt"

// // // go语言允许你控制特定集合的数据结构、分配的数量以及内存访问模式，这些对构建良好的系统是非常重要的：指针对性能的影响是不言而喻的，而如果你想要做
// // // 的是系统编程、操作系统或者网络应用，指针更是不可或缺的一部分
// // func main() {
// // 	// 下面的代码，输出的内存地址，每次都会变化
// // 	var i1 = 5
// // 	fmt.Printf("An integer: %d, it's location in memory:: %p\n", i1, &i1)

// // 	// 这个地址可以存储在一个叫做指针的特殊数据类型中（说明指针是一种 特殊的数据类型），在本例中是一个指向int的指针i1,此处使用*int表示，如果我们想调用指针P，可以这样声明它
// // 	// var intP *int
// // 	// 然后使用intP = &i1是合法的， 此时intP指向i1 (指针的格式化标示符是 %p)
// // 	// intP存储了i1 的内存地址；它指向了i1 的位置，它引用了i1 的变量
// // 	// 一个指针变量可以指向任何一个值得内存地址，它指向的那个值的内存地址，在32位机器上占用4个字节，在64位机器上占用8个字节，并且与它所窒息那个的值的大小无关。
// // 	// 一个指针变量通常缩写为ptr

// // 	var intP *int //声明指针变量
// // 	intP = &i1    // 获取指针的引用
// // 	fmt.Printf("The value at memory location %p is %d \n", intP, *intP)
// // }
// package main

// import "fmt"

// func main() {
// 	s := "good bye"
// 	var p *string = &s
// 	*p = "ciao"

// 	fmt.Printf("Here is the pointer p: %p\n", p)
// 	fmt.Printf("Here is the string *p: %s\n", *p)
// 	fmt.Printf("Here is the string s: %s\n", s)

// 	// 不能得到一个文字或者常量的地址
// 	const i = 5
// 	ptr := &i
// 	ptr2 := &10

// 	fmt.Printf("can't xxx address: %s", ptr)
// 	fmt.Printf("can't xxx address: %s", ptr2)

// 	// 指针的一个高级应用是你可以传递一个变量的引用（如函数的参数）， 这样不会传递变量的拷贝。指针传递是很廉价的，只占用4或者8个字节
// 	// 当程序在工作中需要占用大量的内存，直到没有任何指针指向它们，所以从它们被创建开始就具有相互独立的生命周期。

// 	// 指针也可以指向另一个指针，并且可以进行任意深度的嵌套，导致你可以有多级的间接引用，但在大多数情况下，这会使得你的代码结构不清晰

// }

// 对于一个空指针的反向引用是不合法的，并且会使程序崩溃
package main

func main() {
	var p *int = nil
	*p = 0
}
