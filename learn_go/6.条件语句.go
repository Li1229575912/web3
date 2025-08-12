package main

import (
	"fmt"
	"time"

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
	/* 条件语句 */
	const (
		A = "Alice"
		B = "Bob"
		C = "Cirs"
	)
	//D := A
	var D string
	switch D {
	case A:
		fmt.Println("this is Alice")
		//fallthrough 强制执行下一句，默认自带break
	case B:
		fmt.Println("this is Bob")
	case C:
		fmt.Println("this is Cirs")
	default:
		fmt.Println("I dont know")
	}

	var x interface{} //判断某个接口中变量类型
	switch i := x.(type) {
	case nil:
		fmt.Printf("x。的类型 %T", i)
	default:
		fmt.Println("x。的类型UNKNOW")
	}

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for true {
			ch1 <- "from 1"
		}
	}()

	go func() {
		for true {
			ch2 <- "from 2"
		}
	}()

	for {
		time.Sleep(1 * time.Second)
		select {
		case s1 := <-ch1:
			fmt.Println(s1)
		case s2 := <-ch2:
			fmt.Println(s2)
		default:
			fmt.Println("no message") //如果两个通道都无数据进入这里
		}
	}
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
