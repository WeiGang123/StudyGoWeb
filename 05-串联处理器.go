package main

import (
	"fmt"
	"net/http"
)

// 可以将处理器函数串联来实现一个请求有多个功能对其进行处理

func handleHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello!\n")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	// 接收一个HandlerFunc 返回一个HandlerFunc
	// 给中间可以加上自己想实现的效果
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "log is here!\n")
		h(writer, request)
	}
}

// 如果是串联处理器也是一样的
type HelloHandle struct {}

func (h HelloHandle) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "hello!\n")
}

// 接收和返回Handler类型
func log1(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "log is here!\n")
		h.ServeHTTP(w, r)
	})
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 串联处理器函数
	// http.HandleFunc("/hello", log(handleHello))
	// 串联处理器
	h := HelloHandle{}
	http.Handle("/hello", log1(h))

	server.ListenAndServe()
}