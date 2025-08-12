package main

type Books struct {
	Title  string
	Author string
	Price  float32
}

type Address struct {
	Street string
}

type Employee struct {
	Name string
	Age  int
	Address
}

type AS struct {
	X int
}
type BS struct {
	Y int
}
