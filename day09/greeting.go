// package main

// // Go不允许函数重载，这将导致一个错误，go不支持这个特性的主要原因是函数重载需要进行多余的类型匹配影响性能；
// // 没有重载意味着只是一个简单的函数调度，所以你需要给不同的函数不同的名字。
// // Go也不支持范型，所以在必要的情况下给不同的类型写单独的函数，这样也可以提升代码可读性
// func main() {
// 	println("In main before calling greeting")
// 	greeting()
// 	println("In main after calling greeting")
// }

// func greeting() {
// 	println("In greeting: Hi!!")
// }

// // 函数也可以以申明的方式被使用，作为一个函数类型
// type binOp func(int, int) int
package main

func main() {}
	Pop()
}

func (st *Stack) Pop() int {
	v := 0

	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}
}
