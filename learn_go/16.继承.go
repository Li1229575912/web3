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

type SpeakerInterface interface {
	Speak()
}

type Animal struct {
	Name string
}

func (a *Animal) Speak() {
	fmt.Println(a.Name, "say hello.")
}

type Dog struct {
	Animal
	Breed string
}

func main() {
	/* 继承 */
	var dog Dog = Dog{Animal: Animal{Name: "bitch"}, Breed: "huhu"}
	dog.Speak()
	fmt.Println(dog.Breed)

	var sperker SpeakerInterface
	sperker = &dog
	sperker.Speak()
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked.")
	return mainName
}
