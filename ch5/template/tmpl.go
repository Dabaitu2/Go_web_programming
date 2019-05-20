package main

import (
	"html/template"
	"net/http"
)

// 根目录是以src来决定的
func process(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("./ch5/template/tmpl.html")
	t.Execute(w, "Hello world")
}

func main()  {
	server := http.Server{
	    Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
