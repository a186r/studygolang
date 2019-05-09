//指针方法和值方法都可以在指针或者非指针上被调用
// 比如下面，List在值上有一个方法Len()，在指针上又一个方法Append()，
// 但是可以看到两个方法都可以在两种类型的变量上被调用。
package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

func main() {
	// 值
	var lst List
	lst.Append(1)
	fmt.Println(lst, lst.Len())

	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Println(plst, plst.Len())
}
