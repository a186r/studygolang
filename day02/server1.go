// go语言的内置库使得写一个类似fetch的web服务器变得异常的简单
// 本节会展示一个微型服务器，返回当前用户正在访问的URL

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		main函数将所有发送到/路径下的请求和handler函数关联起来
		/开头的请求其实就是所有发送到当前站点上的请求，服务监听8000端口
	*/

	http.HandleFunc("/", handler) // 每个请求调用处理程序
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
