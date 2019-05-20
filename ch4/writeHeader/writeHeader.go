package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>GoWebProgramming</title>
</head>
<body>
</body>
</html>`
	w.Write([]byte(str))
}

// 写响应码
func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

// 设置头信息
func HeaderExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "tomoko",
		Threads: []string{"frist", "second", "Third"},
	}
	json, _ := json2.Marshal(post)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeHeader", writeHeaderExample)
	http.HandleFunc("/header", HeaderExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
