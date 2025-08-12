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

func main() {
	/*指针*/
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

	fmt.Println(&numArray1[1])

	var int_ptr *int //空指针 nil
	fmt.Println(int_ptr)
	阿 := 3
	int_ptr = &阿
	fmt.Println(*int_ptr)

	if int_ptr != nil {
		fmt.Println("不是空指针")
	}

	const MAX int = 3

	var intArr = [MAX]int{10, 20, 30}
	fmt.Println(intArr)

	var intArr_p [MAX]*int
	intArr_p[0] = &intArr[0]
	intArr_p[1] = &intArr[1]
	intArr_p[2] = &intArr[2]
	for i := 0; i < MAX; i++ {
		fmt.Println(*intArr_p[i])
	}

	var intArr_pp [MAX]**int
	intArr_pp[0] = &intArr_p[0]
	fmt.Println(**intArr_pp[0])

	var intArr_ppp [MAX]***int
	intArr_ppp[0] = &intArr_pp[0]
	fmt.Println(***intArr_ppp[0])

	var vol, cel float32
	getBatteryInfo(&vol, &cel)
	fmt.Println(vol, cel)

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
