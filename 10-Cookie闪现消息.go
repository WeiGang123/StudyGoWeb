package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

// 闪现消息是 在某种条件满足才在页面上出现一种临时的消息，用户在刷新页面之后将不会看到这个消息
// 一般cookie的值没有包含如空格百分号等特殊字符的话 可以不用编码
// 但是闪现消息本身通常会包含诸如空格这样的特殊字符，所以要进行编码
func setMessage(w http.ResponseWriter, r *http.Request)  {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name: "flash",
		// cookie的value用Base64URL编码
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request)  {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No Message found")
		}
	} else {
		// 设置一个刷新以后就会消失的cookie
		rc := http.Cookie{
			Name: "flash",  // 创建一个同名的cookie
			MaxAge: -1,    // 将MaxAge设置为负数
			Expires: time.Unix(1, 0), // 将Expires设置成一个已经过去的时间。
		}
		// 因为是同名cookie 所以会将原来的cookie给抵消
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w,string(val))
	}
}

func main()  {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
