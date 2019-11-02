package main

import (
	"fmt"
	"net/http"
)

// 为了使不同RUL请求返回不同页面，我们不再在Server结构的Handler字段中指定处理器
// 而是让服务器使用默认的DefaultServeMux。然后通过http.Handle函数来将处理器
// 绑定到DefaultServeMux上去。

type HelloHandler struct {}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	hello := HelloHandler{}
	world := WorldHandler{}

	http.Handle("/hello", &hello)
	http.Handle("/world", &world)

	server.ListenAndServe()
}