package main

import (
	"fmt"
	"net/http"
)

// 处理器函数
func process(w http.ResponseWriter, r *http.Request)  {
	// 第一步 先对请求进行语法分析
	// 进行语法分析可以调用ParseForm方法或者ParseMultipartForm方法
	r.ParseForm() // 对请求进行语法分析

	// 第二步 要根据第一步使用的请求方法来访问相应的Form字段、PostForm字段或MultipartForm字段
	// Form结构是一个映射 它的键是字符串，值是一个由字符串组成的切片 内容总是包含URL中的查询值和form中的表单值
	// 我们可以通过设定一个键让它返回对应的值的字符串切片 如果这个键同时包含了表单值和URL值
	// 那么这些值将会包含在一个切片中 并且表单值在切片中的位置会在URL的值的前面
	fmt.Fprintln(w, r.Form, "\nhello的值为：", r.Form["hello"])

	// 如果只想访问表单键值而不想获取URL中的键值对 我们可以访问Request中的PostForm字段
	fmt.Fprintln(w, r.PostForm,"\nhello的值为：", r.PostForm["hello"])

	// FormValue方法只能获取到给定键的第一个值 PostFormValue只会返回表单键值对，而不会返回URL键值对
	fmt.Fprintln(w, r.FormValue("hello"), "-----", r.PostFormValue("hello"))

}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
