package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}
type NewHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 将字符写回response
	fmt.Fprintf(w, "Hello World!")
}

func (h *NewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 将字符写回response
	fmt.Fprintf(w, "Hello World again!")
}

func main() {
	handler := MyHandler{}
	another := NewHandler{}
	// 生成Server对象配置
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 使用默认的Servermux多路复用器
	http.Handle("/hello", &handler)
	http.Handle("/hi", &another)
	server.ListenAndServe()
}
