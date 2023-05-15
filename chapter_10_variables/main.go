package main

import "fmt"

var pL = fmt.Println

func main() {
	var strVar string
	strVar = "Hello"

	var intVar int = 99

	var testVar = false

	var varOne, varTwo = 1, 2

	autoVar := 4.44
	autoVar = 88.7

	pL(strVar, intVar, testVar, varOne, varTwo, autoVar)
}
