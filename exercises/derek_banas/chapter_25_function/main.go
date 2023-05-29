package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("getOne:", getOne(1))

	a, b := getTwo(1)
	fmt.Println("getTwo:", a, b)

	c, err := withError(4, 2)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println("withError:", c)

	fmt.Println("unlimitedParams:", unlimitedParams(1, 2, 3, 4, 5, 6, 7, 8, 9))

	arrInt := []int{1, 2, 3, 4, 5}
	fmt.Println("The sum is", pointerArray(&arrInt))
	fmt.Println("The array", arrInt)

	var d int = 2
	fmt.Println("Number :", d)
	pointerVariable(&d)
	fmt.Println("Modified to:", d)

	fmt.Println(returnWithPointer(1, 2))
}

func getOne(n int) int {
	return n + 1
}

func getTwo(n int) (int, int) {
	return n + 1, n + 2
}

func withError(a int, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("can't divide by zero")
	}

	return a / b, nil
}

func unlimitedParams(numbers ...int) int {
	var sum int

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func pointerArray(arr *[]int) int {
	var sum int
	for i, number := range *arr {
		sum += number
		(*arr)[i]++
	}

	return sum
}

func pointerVariable(n *int) {
	*n++
}

func returnWithPointer(a int, b int) *int {
	var result int = a + b
	return &result
}
