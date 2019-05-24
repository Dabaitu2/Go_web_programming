package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// 通过ioutil向文件写入字节流
	data := []byte("Hello World!\n")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	// 读取文件字节流
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	// 通过创建文件的方式写入数据
	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
