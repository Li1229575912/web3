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

func get_result() (int, int, string) {
	e, f, g := 1, 2, "wang"
	return e, f, g
}

type Gender string //类型别名
const (
	Male   Gender = "男"
	Female Gender = "女"
)

func GenderSelect(sex Gender) string {
	switch sex {
	case Male:
		return "Male"

	case Female:
		return "Female"
	}
	return "default"
}

func main() {
	/* 运算符 */
	// 有a == 20，b == 30
	a, b := 20, 31
	var c int
	c = a + b
	fmt.Printf("a + b = %d\n", c)
	c = a - b
	fmt.Printf("a - b = %d\n", c)
	c = a * b
	fmt.Printf("a * b = %d\n", c)
	c = b / a
	fmt.Printf("b / a = %d\n", c)
	c = b % a
	fmt.Printf("b %% a = %d\n", c)
	a++
	b--
	fmt.Printf("a++ = %d,b-- = %d\n", a, b)

	if a == b {
		fmt.Println("a == b")
	} else if a < b {
		fmt.Println("a < b")
	} else if a > b {
		fmt.Println("a > b")
	}

	c += a
	fmt.Println(c)

	A := 60
	B := 13
	C := A & B
	fmt.Println(C)

	var c_ptr *int = &C
	fmt.Println(c_ptr)
	fmt.Println(*c_ptr)

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
