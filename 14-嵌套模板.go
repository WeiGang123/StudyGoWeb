package main

import (
	"html/template"
	"net/http"
)

func include(w http.ResponseWriter, r *http.Request) {
	// 如果要嵌套模板需要给ParseFiles传入所有的模板文件并且要求第一个是主模板
	t, _ := template.ParseFiles("template/14-1-main.html", "template/14-2-t.html")
	t.Execute(w, "Hello!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/include", include)

	server.ListenAndServe()
}
