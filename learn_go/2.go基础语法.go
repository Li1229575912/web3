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

func main() {
	/* Go语言基础语法 */
	fmt.Println("Hello" + "," + "World")
	// Sprintf格式化输入
	stockcode := 123
	enddate := "2020-12-31"
	url := "Code = %d , endDate = %s"
	target_url := fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	// Printf格式化输出
	fmt.Printf(url, stockcode, enddate)
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
