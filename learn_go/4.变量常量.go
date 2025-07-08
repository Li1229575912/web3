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
	/* 变量常量 */
	var string_arr []string = []string{"Hi", "Baga"}
	fmt.Println(string_arr)

	var a, b, c uint8 = 1, 2, 3

	d, e, f, g := 4, 5, 6, "lili"
	_, no2, no3 := get_result() //舍弃第一个
	fmt.Println(a, b, c, d, e, f, g, no2, no3)

	const LENGTH uint8 = 3
	const WIDTH uint8 = 5
	area := LENGTH * WIDTH
	fmt.Printf("面积是：%d\n", area)

	const ( //const 用作枚举
		Default = 0
		man     = 1
		woman   = 2
		woerma  = iota
		x
		y
		z = iota //自动计数
	)
	fmt.Println(Default, man, woman, woerma, x, y, z)

	const a1, a2, a3 string = "li", "wa", "ba"
	const b1, b2 = 3, false
	fmt.Println(a1, a2, a3, b1, b2)
	fmt.Println(GenderSelect(Female))
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
