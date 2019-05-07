// 使用自定义包中的结构体
package main

import (
	"fmt"

	"./structPack/structPack"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.M1i = 10
	struct1.Mf1 = 16.

	fmt.Printf()
}
