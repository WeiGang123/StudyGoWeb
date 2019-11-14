package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // 导入数据库驱动
)

type post struct {
	Id      int
	Content string
	Author  string
}

// 定义一个指向sql.DB结构的指针
// 这个结构是一个数据库的句柄，它代表一个包含了零个或任意多个数据库连接的连接池，这个连接池由sql包管理
// 这个连接池是自动管理的，所以尽管我们可以手动关闭连接，但是并不需要这么做。
// 我们在这里定义了一个全局变量，除此之外还可以在创建sql.DB结构后通过传参的方式使用它
var Db *sql.DB

// init函数用来初始化Db这个指针变量
// Go语言的每个包都会自动调用定义在包内的init()函数
func init() {
	var err error
	// 程序可以通过调用Open函数，并将相应的数据库驱动名字以及数据源名字传递给该函数来建立与数据库的连接
	// Open函数第一个参数是告诉程序使用哪个数据库驱动，第二个参数是一个字符串，它会告诉驱动应该如何与数据库进行连接
	// Open函数会返回一个指向sql.DB结构的指针
	// Open函数真正的作用是设置连接数据库所需的各个结构，并以惰性的方式，等到真正需要的时候才会建立相应的数据库连接
	// 所以说它不会立即的建立连接，它在执行时候并不会真正地与数据库进行连接
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func posts(limit int) (posts1 []post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		myPost := post{}
		err = rows.Scan(&myPost.Id, &myPost.Content, &myPost.Author)
		if err != nil {
			return
		}
		posts1 = append(posts1, myPost)
	}
	rows.Close()
	return
}

// 通过向函数传递id来返回一个包含完整的数据的变量
func getPost(id int) (myPost post, err error) {
	// 先创建一个空结构体用来存储数据
	myPost = post{}
	// 利用QueryRow进行查询并使用Scan来进行数据的写入。
	// 我们直接使用sql.DB的QueryRow而没有使用sql.Stmt的 是因为我们无需重复的执行sql语句
	// 其实两种都可以使用，这里只是展示了另一种写法。
	err = Db.QueryRow("select id, content, author from posts where id"+
		" = $1", id).Scan(&myPost.Id, &myPost.Content, &myPost.Author)
	return
}

func (myPost *post) Create() (err error) {
	// 定义一条SQL预处理语句，这就是一个SQL语句模板，在执行时需要给语句中的参数提供具体值
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	// Prepare创建预处理语句
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// QueryRow方法用来调用预处理语句，并把来自接收者(mypost)的数据传递给该方法
	// 如果被执行的sql语句返回多于一个sql.Row，这个方法只能返回第一个sql.Row，并丢掉其他的
	// Scan会将查询到的行中的值复制到程序为其提供的参数里面，在这里会把id列的值设置为接收者的值
	err = stmt.QueryRow(myPost.Content, myPost.Author).Scan(&myPost.Id)
	return
}

func (myPost *post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 "+
		"where id = $1", myPost.Id, myPost.Content, myPost.Author)
	return
}

func (myPost *post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", myPost.Id)
	return
}

func main() {
	myPost := post{Content: "Hello World!", Author: "Sau sheong"}

	fmt.Println(myPost)
	myPost.Create()
	fmt.Println(myPost)

	readPost, _ := getPost(myPost.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	myPosts, _ := posts(4)
	fmt.Println(myPosts)

	readPost.Delete()
}
