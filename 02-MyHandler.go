package main

import (
	"fmt"
	"net/http"
)

// 编写一个处理器来代替默认的多路复用器来处理请求

// 所谓处理器， 就是只要有个接口，它拥有一个ServeHTTP方法，
// 并且该方法的签名(参数)和以下的一样，也就是这个参数是固定的。
// ServeHTTP(http.ResponseWriter, *http.Request)
// 那么这个接口就是一个处理器。

// 此代码利用一个处理器去处理请求， 结果是不同URL请求返回的都是相同的内容

type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello World!")
}

func main()  {
	handle := MyHandler{} // 定义一个处理器变量
	server := http.Server { // 通过Server结构对服务器进行配置
		Addr: "127.0.0.1:8080",
		Handler: &handle,
	}
	server.ListenAndServe() // 调用Server结构的ListenAndServe方法来启动服务器
}