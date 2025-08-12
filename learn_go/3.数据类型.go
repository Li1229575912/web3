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
	/* 数据类型 */
	var a uint8 = 0xF
	var b uint8 = 15
	var float1 float32 = 3.1415
	var float2 float64 = 10.0
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(float1)
	fmt.Println(float2)
	var float3 float32 = float32(float2)
	fmt.Println(float3)

	var complex_z complex64 = 1.10 + 0.1i
	fmt.Println(real(complex_z), imag(complex_z)) //分别打印实部 虚部

	complex_y := complex(2.1, 1.34) // 另一种写法
	fmt.Println(real(complex_y), imag(complex_y))

	var rune1 rune = 'H'
	var rune2 rune = 'i'
	var rune3 rune = '世'
	var rune4 rune = '界'
	fmt.Println(rune1, rune2, rune3, rune4)
	string1 := string(rune1) + string(rune2) + string(rune3) + string(rune4) //切片转字符串
	fmt.Println(string1)
	var runes []rune = []rune(string1) //字符串转化切片
	fmt.Println(runes)

	var bytes []byte = []byte(string1) //字符串转[]byte切片
	fmt.Println(bytes)
	var bytes2 []byte = []byte{72, 105, 228, 184, 150, 231, 149, 140}
	fmt.Println(string(bytes2)) //[]byte切片转字符串

	fmt.Println(len(string1))       //8
	fmt.Println(len(runes))         //4 rune的长度是按字符计数的
	fmt.Println(len(bytes))         //8
	fmt.Println(string(runes[0:3])) //截取字符串
	var bool1 bool = false
	if !bool1 {
		fmt.Println("bool1 is not true.")
	}

	var string_nil string
	var float4 float32
	var bool2 bool
	fmt.Println(string_nil, float4, bool2) //初始化默认0值

}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
