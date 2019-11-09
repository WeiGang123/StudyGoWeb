package main

import (
	"html/template"
	"net/http"
)

// 模板中的一个参数就是模板中的一个值，常用一个.来表示服务器向模板引擎传递的数据
// 还可以在动作中设置变量，以美元符号$开头

func argu(w http.ResponseWriter, r *http.Request) {
	// 我们向模板传入一个map来体现变量的作用
	m := map[string]string{
		"Addr":    "127.0.0.1:8080",
		"Handler": "nil",
	}
	// 对模板进行语法分析
	t, _ := template.ParseFiles("template/15-arg.html")
	// 执行模板
	t.Execute(w, m)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/argument", argu)
	server.ListenAndServe()
}
