## 题目1：
只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
```go
	var numArray []int = []int{1, 1, 2, 2, 3, 4, 4, 5, 5, 6, 6, 99, 99, 43, 43, 12, 12}

	countMap := make(map[int]int)

	// 第一次遍历：统计次数
	for _, num := range numArray {
		countMap[num]++
	}

	// 第二次遍历：找出次数为1的元素
	for num, count := range countMap {
		if count == 1 {
			fmt.Println(num)
		}
	}
```

## 题目2：
考察：数字操作、条件判断
题目：判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，12321是回文数，而12345不是。

```go
func isPalindromeNumber(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	// 注意：x 为奇数位时需去掉中间位
	return x == reversed || x == reversed/10
}
```

## 题目3：
有效的括号 
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
```go
func isValid(s string) bool {
	stack := []rune{}
	matching := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack = append(stack, ch)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != matching[ch] {
				return false
			}
			stack = stack[:len(stack)-1] // 出栈
		}
	}
	return len(stack) == 0
}
```
## 题目4：
考察：字符串处理、循环嵌套
题目：查找字符串数组中的最长公共前缀
```go
func longestCommonPrefix(strs []string) string {
	//strs[0]
	common_num := len(strs[0])
	for i, s := range strs {
		fmt.Println(i, s)
		for j, c := range s {
			//fmt.Println(j, string(c))
			if c != rune(strs[0][j]) {
				if common_num > j {
					common_num = j //把比对过程中相同数最小值存下
				}
				break
			}
		}
	}
	return strs[0][:common_num]
}
```
## 题目5：
考察：数组操作、进位处理
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
```go
func AddOne(numArr []uint) []uint {
	for i, v := range numArr {
		numArr[i] = v + 1
	}
	return numArr
}
```
## 题目6：
删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
```go
func removeDuplicates(nums []uint) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
```
## 题目7：
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
```go
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 1. 先按起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	// 2. 遍历并合并
	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		current := intervals[i]

		if last[1] >= current[0] {
			// 有重叠：更新右边界
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 没重叠，直接加入
			result = append(result, current)
		}
	}

	return result
}
```
## 题目8：
两数之和 
考察：数组遍历、map使用
题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
```go
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // 数值 → 索引

	for i, num := range nums {
		complement := target - num
		if idx, ok := m[complement]; ok {
			return []int{idx, i}
		}
		m[num] = i
	}
	return nil
}
```
## 题目9：
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
```go
func numAdd(input *int) {
	*input = *input + 10
}
```
## 题目10：
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
```go
func squareSlice(numSlice *[]int) {
	for i, num := range *numSlice {
		(*numSlice)[i] = num * 2
	}
}
```
## 题目11：
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
```go
func jishu() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("奇数", i)
		}
	}
}

func oushu() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数", i)
		}
	}
}


go jishu()
go oushu()
time.Sleep(2 * time.Second) // 等待协程执行完毕
```
## 题目12：
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
```go
func main() {
	fmt.Println("Hello, World!")
	// 定义一些任务
	tasks := []func(){
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务A完成")
		},
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务B完成")
		},
		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("任务C完成")
		},
	}

	runTasks(tasks)
}

// 任务调度器函数，接收一组任务
func runTasks(tasks []func()) {
	var wg sync.WaitGroup

	for i, task := range tasks {
		wg.Add(1)

		go func(index int, t func()) {
			defer wg.Done()

			start := time.Now()
			t() // 执行任务
			duration := time.Since(start)

			fmt.Printf("任务 %d 执行完成，用时：%v\n", index, duration)
		}(i, task) //(i, task) 是传给匿名函数 func(index int, t func()) 的参数，就像调用普通函数一样。
	}

	wg.Wait()
}
```

## 题目13：
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
```go
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

func main() {
	fmt.Println("Hello, World!")
	rectangle := Rectangle{10, 5}
	fmt.Println("Rectangle area:", rectangle.Area())
	fmt.Println("Rectangle perimeter:", rectangle.Perimeter())
	circle := Circle{3}
	fmt.Println("Circle area:", circle.Area())
	fmt.Println("Circle perimeter:", circle.Perimeter())
}
```
## 题目14：
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
```go
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}

func main() {
	fmt.Println("Hello, World!")
	yuangong := Employee{Person{"Yuangong", 25}, 1001}
	yuangong.PrintInfo()
}
```
## 题目15：
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
```go
func sender(chs chan int) {
	fmt.Println("sender start")

	for i := 0; i < 10; i++ {
		chs <- i
		time.Sleep(1 * time.Second)
	}
	close(chs)
	fmt.Println("sender end")
}

func receiver(chs chan int) {
	fmt.Println("receiver start")

	go func() {
		for {
			select {
			case v := <-chs:
				fmt.Println(v)
			default:
				//fmt.Println("没有数据，休息一下")
				time.Sleep(200 * time.Millisecond) // 防止 CPU 爆炸
			}
		}
	}()
}
```
## 题目16：
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
```go
func main() {
	ch := make(chan int, 10) // 带缓冲的通道，容量为10
	var wg sync.WaitGroup

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			ch <- i // 向通道发送数据
			fmt.Printf("生产者发送：%d\n", i)
		}
		close(ch) // 发送完毕后关闭通道
	}()

	// 消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range ch { // 通道关闭后自动退出
			fmt.Printf("消费者接收：%d\n", v)
		}
	}()

	wg.Wait() // 等待生产者和消费者都结束
	fmt.Println("所有数据处理完毕！")
}
```
## 题目17：
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex       // 互斥锁
	var wg sync.WaitGroup   // 等待所有协程完成

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()      // 加锁
				counter++
				mu.Unlock()    // 解锁
			}
		}()
	}

	wg.Wait()
	fmt.Println("最终计数器值：", counter) // 应该是 10000
}
```
## 题目18：
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64       // 必须是 int64 或 uint64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1) // 原子加1
			}
		}()
	}

	wg.Wait()

	// 原子读取最终值（虽然一般打印前可省略，但标准做法是用 Load）
	final := atomic.LoadInt64(&counter)
	fmt.Println("最终计数器值：", final)
}
```
## 题目19：
