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
	/* slice */

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

	var slice1 []Books = make([]Books, 0)
	slice1 = append(slice1, Book1)
	slice1 = append(slice1, Book2)
	fmt.Println(slice1)

	s := []int{1, 2, 3}
	fmt.Println(s)
	var numbers []int = make([]int, 3, 5)
	numbers[0] = 1
	numbers[1] = 3
	numbers = append(numbers, 7)
	numbers = append(numbers, 9)
	numbers = append(numbers, 10)
	numbers = append(numbers, 13)
	fmt.Println(len(numbers), cap(numbers), numbers)

	var numbers1 []int = make([]int, 10, 20)
	copy(numbers1, numbers)
	fmt.Println(numbers1)
	fmt.Println(numbers1[2:5])
	fmt.Println(numbers1[:6])
	fmt.Println(numbers1[4:])

	var slice2 []float32
	if slice2 == nil {
		fmt.Println("切片是空的")
	}

	for i, v := range numbers1 {
		fmt.Println(i, v)
	}

	for i, c := range "hello" {
		fmt.Println(i, string(c))
	}

	channel := make(chan int, 2)
	channel <- 1
	channel <- 2
	close(channel)

	for v := range channel {
		fmt.Println(v)
	}

	map1 := make(map[string]int, 10)
	fmt.Println(map1)
	map1 = map[string]int{
		"apple":  1,
		"orange": 2,
		"banana": 3,
	}
	for key, v := range map1 {
		fmt.Println(key, v)
	}
	//修改键值
	map1["banana"] = 5
	//获取键值
	v1, ok := map1["banana"]
	if ok {
		fmt.Println(v1)
	}
	//获取map长度
	len := len(map1)
	fmt.Println(len)
	//删除元素
	delete(map1, "banana")
	fmt.Println(map1)

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
