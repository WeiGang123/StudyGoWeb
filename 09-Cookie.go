package main

import (
	"fmt"
	"net/http"
)

// 向响应中添加cookie
func setCookie(w http.ResponseWriter, r *http.Request)  {
	// 在go语言中 cookie用Cookie结构表示。
	// 通过Cookie结构生成两个Cookie变量
	c1 := http.Cookie{
		Name:		"first_cookie",
		Value:		"Go Web Programming",
		HttpOnly:	true,
		// HttpOnly是Cookie的扩展功能， 设置为true可以使js脚本无法获得Cookie
		// 其主要目的是为了防止跨站脚本攻击(XSS)对Cookie的信息窃取
	}

	c2 := http.Cookie{
		Name:	"second_cookie",
		Value: "Weigang Study Go",
		HttpOnly: true,
	}

	/*
	// 首先使用Set方法添加第一个cookie，然后再使用Add方法添加第二个cookie
	// Cookie结构中的String方法可以返回一个经过序列化处理的cookie
	// http首部字段的Set-Cookie字段的值就是这些序列化的cookie组成的
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	 */

	// 除了上面的通过修改首部字段的方法添加cookie，go语言还提供了一个cookie方法
	// 需要注意的是SetCookie方法的参数传递的是Cookie结构的指针。
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

// 从浏览器获取cookie
func getCookie(w http.ResponseWriter, r *http.Request)  {
	/*
	// 第一种方法通过Header返回的首部的映射中将键为Cookie的值取出
	// 结构是一个切片，切片中是一个字符串，而字符串中又包含了多个cookie
	// 如果取得单独的cookie键值对，就得对字符串进行语法分析
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
	*/

	// 还有一种方法可以更容易的获取cookie
	// 通过Request结构的Cookie方法可以获取某个指定名字的cookie
	// 如果不存在，则返回一个错误
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first_cookie")
	}
	// Cookie方法只能返回一个，要想返回多个可以使用Cookies方法
	// 这个方法返回的结构和通过Header获取的结构是一样的。
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	server.ListenAndServe()
}
