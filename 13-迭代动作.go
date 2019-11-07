package main

import (
	"html/template"
	"net/http"
)

// 迭代动作能向模板传入一个切片元素, 在模板里面利用 range就可以遍历出切片中的所有元素
// 迭代循环中的{{.}} 表示的是当前被迭代的切片元素.
func iteration(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/13-iteration.html")
	name := []string{"wang", "li", "zhang", "liu"}
	t.Execute(w, name)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/iteration", iteration)

	server.ListenAndServe()
}
