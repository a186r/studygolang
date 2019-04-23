// // // package main

// // // import (
// // // 	"fmt"
// // // 	"strings"
// // // )

// // // // Join用于将元素类型为string的slice使用分隔符来拼接成一个字符串
// // // // strings.Join(sl []string, sep string) string
// // // func main() {
// // // 	str := "The quick brown fox jums over the lazy dog"
// // // 	sl := strings.Fields(str)
// // // 	fmt.Printf("Splitted in slice: %v\n", sl)
// // // 	for _, val := range sl {
// // // 		fmt.Printf("%s -", val)
// // // 	}

// // // 	fmt.Println()

// // // 	str2 := "GO1| The ABC of Go|25"
// // // 	sl2 := strings.Split(str2, "|")
// // // 	fmt.Printf("Splitted in slice: %v \n", sl2)
// // // 	for _, val := range sl2 {
// // // 		fmt.Printf("%s -", val)
// // // 	}
// // // 	fmt.Println()

// // // 	str3 := strings.Join(sl2, ";")
// // // 	fmt.Printf("sl2 joined by ;: %s\n", str3)
// // // }

// // // // 函数strings.NewReader(str)用于生成一个Reader并读取字符串中的内容，然后返回指向该Reader的指针，从其他类型读取内容的函数还有：
// // // // Read() 从byte中读取内容
// // // // ReadByte()和ReadRune()从字符串中读取下一个byte或者rune.
// // // // 字符串与其他类型的转换
// // // // 与字符串相关的转换都是通过strconv包实现的
// // // // 该包包含了一些变量用于获取程序运行的操作平台下的int类型所占的位数

// // package main

// // import (
// // 	"fmt"
// // 	"strconv"
// // )

// // func main() {
// // 	// strconv.Iota(i int) string 返回数字i所表示的字符串类型的十进制数
// // 	fmt.Println(strconv.Itoa(123))
// // 	// strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将64位浮点型的数字转换为字符串，其中fmt代表格式，prec表示精度，bitSize则使用32，表示float32， 64表示float64
// // 	fmt.Println(strconv.FormatFloat(3.1415926535987, 'e', 3, 64)) //'e'表示十进制指数
// // }

// // // 针对从字符串转换为数字类型，go提供了以下函数
// // // strconv.Atoi(s string) (i int, err error) //将字符串转换为int类型
// // // strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为float64类型
// // // 利用多返回值的特性，这些函数会返回2个值，第一个是转换后的结果，第二个是可能出现的错误，因此，我们一般使用以下形式来进行从字符串到其他类型的转换：
// // // var, err = strconv.Atoi(s)
// // strconv.IntSize 用于获取程序运行平台下int类型所占的位数
// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	var orig string = "666"
// 	var an int
// 	var newS string

// 	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

// 	an, _ = strconv.Atoi(orig)
// 	fmt.Printf("The integer is: %d\n", an)
// 	an = an + 5
// 	newS = strconv.Itoa(an)
// 	fmt.Printf("The new string is : %s\n", newS)
// }
