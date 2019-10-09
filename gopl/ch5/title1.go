package ch5

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func title1(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	//	检查Content-Type是HTML
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	time.Sleep(10 * time.Second)
}

//trace记录函数退出的时间，并输出总耗时
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

//func double(x int) int {
//	return x + x
//}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d/n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() {result += x}()
	return double(x)
}

//fmt.Println(triple(4)) // 12

//在循环体中使用defer语句要特别注意，因为defer会在整个循环结束后才会调用，可能会导致很多资源没有关闭，可以使用下面
//的处理方法，在循环体外再写一个方法，在该方法中调用defer

func deferTest(filenames []string) error {
	for _, filename := range filenames{
		if err := doFile(filename); err != nil{
			return err
		}
	}
	return nil
}

//
func doFile(filename string) error{
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer f.Close()
}


