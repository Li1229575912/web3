package main

import (
	"fmt"
	"math"

	_ "github.com/learn/init_order/pkg1"
)

var g int

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

func max(num1, num2 int, name string) (int, string) {
	return num1 + num2, name
}

func swap(x, y *string) {
	z := *x
	*x = *y
	*y = z
}

func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type Circle struct {
	radius float64
}

// 该method属于Circle类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
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
	num, name := max(3, 4, "wu")
	fmt.Println(num, name)

	var sl string = "li"
	var sw string = "wang"
	swap(&sl, &sw)
	fmt.Println(sl, sw)

	getSqrtRoot := func(x, y float64) (float64, float64) {
		return math.Sqrt(x), math.Sqrt(y)
	}
	_, b := getSqrtRoot(9, 16)
	fmt.Println(b)

	for i := 0; i < 3; i++ {
		//next_num := getSequence()
		//fmt.Println(next_num())
		fmt.Println(getSequence()())
	}

	var c1 Circle
	c1.radius = 10.00
	fmt.Println(c1.getArea())

	g = 42783
	fmt.Println(mainName, g)

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
