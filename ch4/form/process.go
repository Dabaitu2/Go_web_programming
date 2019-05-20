package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Fprintln(w, r.Form)
	// 使用postform
	//
	// 可以只获取表单中的元素，而不获取query中的同名元素
	//fmt.Fprintln(w, r.PostForm)
	// 处理formdata使用MultipartForm
	//r.ParseMultipartForm()
	//fmt.Fprintln(w, r.MultipartForm)
	// 使用formValue可以不写parseForm直接获取表单元素
	//fmt.Fprintln(w, r.FormValue("hello"))
	// 获取文件方法
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

}

func betterProcess(w http.ResponseWriter, r *http.Request)  {
	file, _,  err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/betterProcess", betterProcess)
	server.ListenAndServe()
}
