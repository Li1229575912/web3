package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	_ "learn_go2/shabi"
	"os"
	"regexp"
	"sync"
	"unsafe"
)

var (
	xiaoming            string = "xiaoming"
	liliang             string = "liliang"
	name1, name2, name3        = xiaoming, liliang, "xiaohong"
)

const A byte = 'A'
const ni rune = '你'
const hao rune = '好'
const PI = 3.14

// Go中enum直接用const定义
const (
	Monday  = 1
	Tuesday = iota + 1
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	value1, value2 := 1, 2
	v3, v4 := jiaandjian(value1, value2)
	fmt.Println(v3, v4)
	fmt.Println("Hello, world!")
	fmt.Println(name1, name2, name3)
	fmt.Println(A, string(ni)+string(hao))
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	unsafename1ptr := unsafe.Pointer(&name1)
	fmt.Println(unsafename1ptr, *(*string)(unsafename1ptr))
	p2 := unsafe.Pointer(uintptr(unsafename1ptr))
	p3 := (*string)(p2) //unsafe.Pointer用完要指定转回什么类型指针，不然编译器不知道
	*p3 = "baga"
	fmt.Println(name1)

	yuwen := Books{"语文", "老子", 88.88}
	shuxue, yingyu := Books{"数学", "苏轼", 99.99}, Books{"英语", "朱熹", 100.00}
	fmt.Println(yuwen, shuxue, yingyu)
	bookstore(&yingyu)
	yuangong1 := Employee{"yuangong1", 20, Address{"shanghai"}}
	fmt.Println(yuangong1)
	loop()

	// 接口类型转换
	var i interface{} = "hello"
	// 直接断言
	s := i.(string)
	fmt.Println(s) //如果不是string 就会panic
	// 类型断言
	s, ok := i.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println("not string")
	}
	//结构体之间必须显式转换
	a := AS{100}
	b := BS{Y: a.X}
	fmt.Println(a, b)

	Talk(Dog{name: "wangwang"})

	p := Person{name: "xiaoming"}
	p.Runing()
	p.Speak()

	xingwei(&Person{name: "xiaohong"})
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println(errordemo())

	go sayHello()
	go sayGoodbye()
	var wg sync.WaitGroup
	ch1 := make(chan int, 2)
	wg.Add(1)
	go func() {
		select {
		case <-ch1:
			fmt.Println("ch1", <-ch1)
		default:
			fmt.Println("default")
		}
	}()
	go square(ch1, 3, &wg)
	//square_result := <-ch1 // 接收数据
	//fmt.Println(square_result)
	//time.Sleep(2 * time.Second)
	wg.Wait()

	matched, err := regexp.MatchString("^h.*o$", "hello")
	fmt.Println(matched) // true
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re.MatchString("abc123")) // true

	//open file
	f, err := os.Create("new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	f.WriteString("hello, Go\n")

	file, _ := os.OpenFile("new.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	file.WriteString("追加内容\n")

	writer := bufio.NewWriter(file)
	writer.WriteString("bufio写入第1行内容\n")
	writer.WriteString("bufio写入第2行内容\n")
	writer.WriteString("bufio写入第3行内容\n")
	writer.Flush() // 必须刷新缓冲区

	src, _ := os.Open("new.txt")
	dst, _ := os.Create("new2.txt")
	defer dst.Close()
	defer src.Close()
	io.Copy(dst, src) // 复制文件内容

	err1 := os.Remove("new.txt")
	if err1 != nil {
		fmt.Println(err1)
	}
	err2 := os.Remove("new2.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
}

func jiaandjian(a, b int) (int, int) {
	return a + b, a - b
}

func loop() {
	numbers := []int{1, 2, 3, 4, 5}
	for key, value := range numbers {
		fmt.Println(key, value)
	}
}

func bookstore(book *Books) {
	fmt.Println(book.Title, book.Author, book.Price, "xiexieguanglin")
}

// interface调用
func Talk(s Speaker) {
	fmt.Println(s.Speak())
}

func xingwei(b PersonBehaiver) {
	b.Runing()
	b.Speak()
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为零")
	}
	return a / b, nil
}

func sayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello，gogogo")
	}
}

func sayGoodbye() {
	for i := 0; i < 5; i++ {
		fmt.Println("Goodbye，gogogo")
	}
}

func square(out chan int, x int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		out <- x * x // 发送数据
	}

}
