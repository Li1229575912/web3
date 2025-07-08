package main

import (
	"errors"
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

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

func (c Circle) Perimeter() float64 {
	return c.radius * 3.14 * 2
}

func printValue(val interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", val, val)
	_, ok := val.(string)
	if ok {
		fmt.Print("is string \n")
	}
	switch v := val.(type) {
	case int:
	case string:
		fmt.Println(v, "string")
	case bool:
	default:

	}
}

type error interface {
	Error() string
}

type DivideError struct {
	Dividend int
	Divisor  int
}

func (e DivideError) Error() string {
	return fmt.Sprintf("can not divide %d by %d", e.Dividend, e.Divisor)
}

type DivideError2 struct {
	dividee int
	divider int
}

func (de *DivideError2) Error() string {
	strFormat := `
	Cannot proceed,the divider is zero
	dividee:%d
	divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError2{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

var ErrNotFound error = errors.New("not found")

func findItem(id int) error {
	return fmt.Errorf("database error: %w", ErrNotFound)
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	panic("something went wrong")
}

func main() {
	/*接口 */
	c := Circle{radius: 5}
	var s Shape = c
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
	printValue("hi")
	printValue(123)
	err := errors.New("this is an error")
	fmt.Println(err)

	e := DivideError{Dividend: 6, Divisor: 10}
	var err1 error = e
	fmt.Println(err1.Error())

	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10=", result)
	}
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	err3 := findItem(1)
	if errors.Is(err3, ErrNotFound) {
		fmt.Println("Item not found")
	} else {
		fmt.Println("Other error: ", err3)
	}
	safeFunction()
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
