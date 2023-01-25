package main

import "fmt"

func main() {
	var a int = 60
	var b float32
	b = 6.89
	c := "I like eating meat"

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n\n", c, c)

	var conferenceName = "Go Conference"

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("Get your tickets here to attend")

}
