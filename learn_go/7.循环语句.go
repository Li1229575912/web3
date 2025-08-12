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
	/* 循环语句 */
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i >= 5 && i < 8 {
			fmt.Println("loop continue , now num:", i)
			continue
		} else if i >= 8 {
			fmt.Println("loop cancel , now num:", i)
			break
		}
	}

	var a int = 10
	//循环
LOOP:
	for a < 20 {
		if a == 15 {
			//跳过迭代
			a++
			goto LOOP
		}
		fmt.Println("a的值", a)
		a++
	}

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
