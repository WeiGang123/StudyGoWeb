package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")

	// 使用ioutil包中的WriteFile和ReadFile函数进行读写
	// 使用WriteFile函数的时候需要有一个文件名，传入的数据和一个用于设置权限的数字作为参数
	// 传给WriteFile函数和ReadFile函数的返回值都是一个字节组成的切片([]byte)。
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Print(string(read1))

	// 使用File结构进行读写文件。相比上面的方法来说较为麻烦但灵活性较高
	// 首先利用os包的Create函数创建文件，传入文件名。
	file1, _ := os.Create("data2")
	// 使用defer关闭文件，防止之后用完忘记关闭。defer语句在函数结束之前执行。
	defer file1.Close()

	// 使用Write方法写入数据
	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	// 读数据之前先使用os包的Open函数打开文件
	file2, _ := os.Open("data2")
	defer file2.Close()

	// 在使用Read进行读，要先申请一个空间，然后作为参数传入，内容将会存在这个空间中
	// 它会返回读取的字节数和一个错误
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from\n", bytes)
	fmt.Println(string(read2))

}
