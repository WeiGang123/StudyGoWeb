package main

import (
	"fmt"
	"net/http"
)

// 处理器函数就是与处理器有相同行为的函数。
// 也就是说处理器函数和ServeHTTP方法有相同的签名，或者说，
// 它们接受ResponseWriter和指向Request结构的指针作为参数
// 这样的函数称为处理器函数

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintf(w, "World!")
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// HandleFunc 函数 可以将一个带有正确签名的函数f转换成一个
	// 带有方法f的Handler 并与默认的多路复用器DefaultServeMux绑定
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}