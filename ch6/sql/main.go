package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// 大写的才是公开字段
type Post struct {
	Id      int
	Content string
	Author  string
}

// 创建数据库句柄
var Db *sql.DB

func init() {
	var err error
	// 连接到数据库
	Db, err = sql.Open("postgres", "user=tomokokawase dbname=gwp password=zhy677097 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err := rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return posts, err
		}
		posts = append(posts, post)
	}
	rows.Close()
	// 会知道自动return什么的
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// 对象的方法
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("update  posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	// 不需要设置，数据库会自增
	post := Post{Content: "Hello World", Author: "tomokokawase"}
	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "heheheheh"
	readPost.Author = "zhangdaming"
	readPost.Update()

	posts, _ := Posts(1)
	fmt.Println(posts)

	//readPost.Delete()
}
