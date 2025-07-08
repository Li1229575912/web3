package main

import (
	"fmt"
	"strconv"

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

var a int = 0

func recursion() {
	a++
	fmt.Println(a)
	if a < 50 { //递归函数必须设置退出条件
		recursion()
	}
}

// 阶乘 5！= 5*4*3*2*1 = 120
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * (factorial(n - 1))
}

func main() {
	/* 递归函数 */
	recursion()
	fmt.Println(factorial(5))

	a := 3
	b := 5
	var mean float32 = float32(a) / float32(b)
	fmt.Println(mean)

	var str string = "1888"
	num, err := strconv.Atoi(str)
	if err == nil {
		fmt.Println(num)
	}

	num_f, err := strconv.ParseFloat("3.14", 64)
	fmt.Println(num_f)

	str_f := strconv.FormatFloat(6.23, 'f', 2, 64)
	fmt.Println(str_f)

	var int123 int = 1234
	str123 := strconv.Itoa(int123)
	fmt.Println(str123)

	var i interface{} = "Hello World"
	str_i, ok := i.(string)
	if ok {
		fmt.Println(str_i, "is string")
	} else {
		fmt.Println(str_i, "is not string")
	}

	var v interface{}
	switch v := v.(type) {
	case string:
		fmt.Println(v, "is string")
	default:
		fmt.Println("wei zhi")
	}

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
