package main

import (
	"fmt"
	"net/http"
)

func getHeaders(w http.ResponseWriter, r *http.Request)  {
	h := r.Header  // 获取头部信息
	for t := range h {
		fmt.Fprintln(w, t, ":", h[t])
	}
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", getHeaders)
	server.ListenAndServe()
}
