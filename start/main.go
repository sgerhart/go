package main

import "fmt"

// Package Level Variable
var h int = 42

// Variable Grouping
var (
	name = "Steve"
	age  = 46
)

const (
	test = 1 << iota
	test2
	test3
	test4
	test5
	test6
)

func main() {

	/*// Variable declaration without value, value on other line
	var i int
	i = 42
	// Variable declaration on the same line as the number
	var x int = 40
	// 	Short declaration but can only be described in a function
	k := 50
	fmt.Println(i, x, k, h)

	fmt.Println("Hello, World")

	const myConst int = 8

	fmt.Printf("%v, %T", myConst, myConst)*/
	var role byte = test | test2 | test4 | test6
	fmt.Printf("Is test %v\n", role&test5 == test5)

	fmt.Printf("%b", role)
}
