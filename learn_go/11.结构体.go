package main

import (
	"fmt"

	_ "github.com/learn/init_order/pkg1"
)

const mainName string = "main"

var mainVar string = getMainVar()

func init() {
	fmt.Println("main init method invoked.")
}

//求5个数平均值

func getAverage(numArray []int, size uint) int {
	var avg int = 0
	for i := 0; i < int(size); i++ {
		avg += numArray[i]
	}
	return avg / int(size)
}

func getBatteryInfo(vol *float32, cel *float32) {
	*vol = 36.00
	*cel = 4650
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book Books) {
	fmt.Println(book)
}

func main() {
	/* 结构体 */

	// var book1 = Books{"123", "eq", "dhuaf sub", 124}
	// fmt.Println(book1)
	fmt.Println(Books{"123", "eq", "dhuaf sub", 124})
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	printBook(Book1)
	printBook(Book2)

	var book_p [2]*Books
	book_p[0] = &Book1
	book_p[1] = &Book2
	fmt.Println(book_p)
	fmt.Println(*book_p[1])

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
