package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("11-tmpl.html")
	if t != nil {
		t.Execute(w, "Hello World!")
	} else {
		fmt.Fprintln(w, "No html")
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
