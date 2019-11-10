package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// 在之前我们使用包含动作来实现了模板 但这种方式不适合开发复杂的Web应用
// 因为这种方式需要将大量代码硬编码到处理器里面，还需要创建大量的模板文件，无法拥有公用的公共布局
// 为此，我们可以使用定义动作(define action)在模板文件里面显式的定义模板

func layout(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("template/17-layout.html", "template/17-content.html")
	} else {
		// 当我们只分析layout.html文件时就找不到content模板，这时调用的block定义的content模板
		t, _ = template.ParseFiles("template/17-layout.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/layout", layout)

	server.ListenAndServe()
}
