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

func main() {
	/*数组*/
	balance := [10]float32{1.00, 2.00, 3.00}
	fmt.Println(balance)

	var num = [5]int8{-1, 0, 1, 2, 3}
	fmt.Println(num)
	for i := 0; i < 5; i++ {
		fmt.Println(num[i])
	}

	var strArray []string = []string{"Alice", "Bob", "Cris"}
	fmt.Println(strArray)
	var ptrArray []*string = []*string{&strArray[0], &strArray[1]}
	fmt.Println(*ptrArray[0], *ptrArray[1])

	values := [][]int{}
	var row1 = []int{1, 2, 3}
	var row2 = []int{4, 5, 6}

	values = append(values, row1)
	values = append(values, row2)
	fmt.Println(values[0])
	fmt.Println(values[1])

	var numArray1 = []int{1, 2, 3, 4, 5}
	fmt.Println(getAverage(numArray1, 5))

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
