package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main()  {
	/*
	// 最简单的web服务器
	http.ListenAndServe("127.0.0.1:8080", nil)
	*/

	/*
	// 带附加配置的服务器
	server := http.Server {
		Addr: "127.0.0.1:8080", // 配置地址
		Handler: nil, // 配置处理器 为空表示使用默认的多路复用器DefaultServeMux
	}
	server.ListenAndServe()
	 */

	// 自定义处理器函数
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}