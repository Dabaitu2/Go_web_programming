package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct {
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "World!")
}

func main()  {
	h := HelloHandler{}
	w := WorldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	// 实际上是在调用 defaultServerMux的方法
	http.Handle("/hello", &h)
	http.Handle("/world", &w)
	server.ListenAndServe()
}