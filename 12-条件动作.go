package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func condition(w http.ResponseWriter, r *http.Request) {
	// 首先进行语法分析
	t, _ := template.ParseFiles("template/12-condition.html")
	rand.Seed(time.Now().Unix()) // 随机数种子
	// 将数据应用到模板里
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/condition", condition)

	server.ListenAndServe()
}
