package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	fmt.Println("Woof!my name is ", d.name)
	return "Woof!" + d.name
}

type PersonBehaiver interface {
	Speak() string
	Runing()
}

type Person struct {
	name string
}

func (p *Person) Speak() string {
	fmt.Println("Hello! my name is ", p.name)
	return "Hello!" + p.name
}

func (p *Person) Runing() {
	fmt.Println("I am running")
}

type MyError struct {
	Msg  string
	Code int
}

func (e MyError) Error() string {
	return fmt.Sprintf("错误 %d: %s", e.Code, e.Msg)
}

func errordemo() error {
	return MyError{"连接失败", 500}
}
