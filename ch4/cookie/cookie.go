package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request)  {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web programming",
		HttpOnly: false,
	}
	c2 := http.Cookie{
		Name: "Second_cookie",
		Value: "Go Web yeee programming",
		HttpOnly: false,
	}
	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	w.Write([]byte("hello"))
}

func getCookie(w http.ResponseWriter, r *http.Request)  {
	c1, err :=  r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Can not get that cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
