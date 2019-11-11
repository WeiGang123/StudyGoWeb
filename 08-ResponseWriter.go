package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

// 用来序列化json
type Post1 struct {
	User    string
	Threads []string
}

// 通过第一个参数（ResponseWriter接口）来向客户端返回响应
// ResponseWriter接口有三个方法
// - Write 接受一个字节数组作为参数，并将数组中的字节写入HTTP响应的主体中
// - WriteHeader 接受一个代表HTTP响应状态码的整数作为参数，并将这个整数作为HTTP响应的返回状态码
//	 如果用户在调用Write方法之前没有执行WriteHeader，那么程序会默认将200 OK 作为响应的状态码
// - Header 调用此方法可以获得一个由首部组成的映射，修改这个映射就可以修改首部，修改后的首部被包含在HTTP响应里面，发送至客户端

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	// 对HTTP响应添加Location的首部字段进行HTTP的重定向
	w.Header().Set("Location", "http://www.weigang.wang")
	// WriteHeader方法在执行完毕以后就不允许再对首部进行写入了，
	// 所以先对首部进行写入，再修改状态码
	w.WriteHeader(302)
}

// 向客户端直接返回json数据
func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post1{
		User:    "wang weigang",
		Threads: []string{"first", "second", "third"},
	}
	json, _ := json2.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)

	server.ListenAndServe()
}
