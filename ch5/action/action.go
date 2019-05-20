package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./ch5/action/action.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func iterate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./ch5/action/iterate.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func withFunc(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("./ch5/action/with.html")
	t.Execute(w, "hello")
}

func include(w http.ResponseWriter, r *http.Request)  {
	t, _ := template.ParseFiles("./ch5/action/include.html", "./ch5/action/inside.html")
	t.Execute(w, "hello world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/iterate", iterate)
	http.HandleFunc("/with", withFunc)
	http.HandleFunc("/include", include)
	server.ListenAndServe()
}
