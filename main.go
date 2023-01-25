package main

import "fmt"

func main() {
	var a int = 60
	var b float32
	b = 6.89
	c := "I like eating meat"

	fmt.Println("Welcome to our conference booking application")
	fmt.Println("Get your tickets here to attend")

	fmt.Printf("%v, %T", a, a)
	fmt.Printf("%v, %T", b, b)
	fmt.Printf("%v, %T", c, c)
}