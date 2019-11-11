package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post2 struct {
	Id      int
	Content string
	Author  string
}

func main() {
	// 使用os包的Create创建一个csv文件
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	// 所有的数据
	allPosts := []Post2{
		{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	// 我们使用NewWriter函数创建一个新的写入器，并把文件当做参数，将其传递给写入器。
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// 由于程序接下来立即要对写入的posts.csv文件进行读取，而刚刚写入的数据有可能还在缓冲区中
	// 所以我们要调用Flush方法来保证缓冲区中的所有数据都已经被正确的写入文件里面了
	writer.Flush()

	// 读取文件
	// 首先打开文件
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 使用NewReader函数来创建一个读取器
	reader := csv.NewReader(file)
	// 将读取器的FieldsPerRecord字段的值设置为负数
	// 这样即使读取器在读取时发现记录缺少某些字段也不会被中断
	// 如果值为正数，那么这个值就是要求从每条记录里读取出的字段的数量，
	// 如果读取出的字段少于这个数 go就会抛出一个错误
	// 如果值为0，则读取器就会将读取的第一条记录的字段数量用作FieldsPerRecord的值
	reader.FieldsPerRecord = -1
	// 一次性的读取文件中包含的所有记录，如果文件较大，可采用其他方法每次一条的读取文件
	// ReadAll返回一个由切片组成的切片
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post2
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post2{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
