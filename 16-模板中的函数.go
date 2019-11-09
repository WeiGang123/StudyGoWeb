package main

import (
	"html/template"
	"net/http"
	"time"
)

// Go函数也可以作为模板参数。
// 1、首先需要创建一个名为FuncMap的映射，并将映射的键设置为函数的名字，值设置为实际定义的函数
// 2、将FuncMap与模板进行绑定.

// 创建一个函数用来将time对象转换成字符串
func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func myFunc(w http.ResponseWriter, r *http.Request) {
	// 创建映射
	funcMap := template.FuncMap{"fdate": formatDate}
	// 将映射绑定到模板上
	t := template.New("16-func.html").Funcs(funcMap)
	// 对模板进行语法分析
	t, _ = t.ParseFiles("template/16-func.html")
	// 执行模板
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/func", myFunc)

	server.ListenAndServe()
}
