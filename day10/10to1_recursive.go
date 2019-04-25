// 使用递归函数从10打印到1
package main

import "fmt"

func main() {
	printrec(1)
}

func printrec(i int) {
	if i > 10 {
		return
	}

	printrec(i + 1)
	fmt.Printf("%d\n", i)
}
