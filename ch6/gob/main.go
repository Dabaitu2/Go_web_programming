package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

// 空接口就说明随便什么类型都可以传入
func store(data interface{}, filename string) {
	// 建立一个空的可读可写缓冲区
	buffer := new(bytes.Buffer)
	// 将缓冲区装上gob编码器
	encoder := gob.NewEncoder(buffer)
	// 将数据编码后送入缓冲区
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	// 将缓冲区数据以字节流形式送入文件
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	// 读取原始字节流
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// 为字节流创建缓冲区
	buffer := bytes.NewBuffer(raw)
	// 为缓冲区装载解码器
	dec := gob.NewDecoder(buffer)
	// 解码缓冲区数据, 并送入data
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{Id: 1, Content: "Hello World!", Author: "tomokokwase"}
	store(post, "post1")
	var postread Post
	load(&postread, "post1")
	fmt.Println(postread)
}
