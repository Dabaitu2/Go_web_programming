package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	csvFile, err := os.Create("post.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	allPosts := []Post{
		{Id: 1, Content: "Hello World!", Author: "tomoko"},
		{Id: 2, Content: "World Hello!", Author: "kawase"},
		{Id: 3, Content: "World Byeee!", Author: "kawase"},
		{Id: 4, Content: "World Yeahh!", Author: "somebody"},
	}

	// 打开写入流
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		// string()无法强转int，会转成内码, 只有在[]byte中能被识别
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// 刷新输入缓冲区
	writer.Flush()

	file, err := os.Open("post.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
