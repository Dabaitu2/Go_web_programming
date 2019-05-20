package main

import (
	"fmt"
	"net/http"
)


func hello(w http.ResponseWriter, r *http.Request) {
	// 将字符写回response
	fmt.Fprintf(w, "Hello World!")
}

func hi(w http.ResponseWriter, r *http.Request) {
	// 将字符写回response
	fmt.Fprintf(w, "Hello World again!")
}

func main() {
	// 生成Server对象配置
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 使用默认的Servermux多路复用器
	// HandleFunc 不需要初始化符合存在ServeHTTP接口的对象
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hi", hi)
	server.ListenAndServe()
}
