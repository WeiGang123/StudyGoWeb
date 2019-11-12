package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type post struct {
	Id      int
	Content string
	Author  string
}

// 储存数据
func store(data interface{}, filename string) {
	// bytes.Buffer结构是拥有Read方法和Write方法的可变长度字节缓冲区，
	// 它既是读取器又是写入器
	buffer := new(bytes.Buffer)
	// 创建一个gob编码器
	encoder := gob.NewEncoder(buffer)
	// 调用编码器的Encode方法将数据编码进缓冲区里面
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	// 最后将缓冲区中已编码的数据写入文件
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

// 载入数据
func load(data interface{}, filename string) {
	// 先读取文件内容
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	// 根据文件内容来创建一个缓冲区
	buffer := bytes.NewReader(raw)
	// 根据缓冲区创建一个gob编码器
	dec := gob.NewDecoder(buffer)
	// 根据编码器的Decode方法来讲数据载入到给定的post结构里面
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	myPost := post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store(myPost, "post1")
	var postRead post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
