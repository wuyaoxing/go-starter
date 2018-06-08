// 程序所属包
package main

// 导入依赖包
import "fmt"

// 常量定义
const NAME = "name"

// 全局变量的声明与赋值
var a string = "a"

// 类型声明
type numInt int

// 结构声明
type Books struct {
	title string
	author string
	subject string
	book_id int
}

// 接口声明
type Ibook interface {

}

// 函数定义
func printBook( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}

// main
func main () {
	var Book Books

	Book.title = "Go 语言"
	Book.author = "go go"
	Book.subject = "Go 语言教程"
	Book.book_id = 65481407

	fmt.Println("This is base grammar for go.")
	fmt.Println(NAME)
	fmt.Println(a)

	printBook(Book)
}