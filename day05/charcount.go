// 统计输入中每个Unicode码点出现的次数
// 虽然Unicode全部码点的数量巨大，但是出现在特定文档中的字符种类没有多少
// 使用map可以比较自然的跟踪那些出现过的字符的次数
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	// 声明一个变量，表示Unicode字符的数量
	counts := make(map[rune]int)
	// UTF-8编码的长度
	var utflen [utf8.UTFMax + 1]int
	//有效的UTF-8编码的数量
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() //返回rune,nbytes,error
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++

	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
