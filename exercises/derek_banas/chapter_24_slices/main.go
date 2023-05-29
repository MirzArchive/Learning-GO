package main

import "fmt"

func main() {
	arrStr := make([]string, 6)

	arrStr[0] = "Cornelius"
	arrStr[1] = "Dawg"
	arrStr[2] = "Nusnus"
	arrStr[3] = "Viton"
	arrStr[4] = "Swington"
	arrStr[5] = "Jejctment"

	for i, v := range arrStr {
		fmt.Printf("index of %d : %s\n", i, v)
	}

	str := arrStr[4][0:2]
	fmt.Println(str)
	str = "BOB"

	fmt.Println(arrStr[4])
	fmt.Println(str)

	arrInt := make([]int, 5)

	for i := range arrInt {
		arrInt[i] = i + 1
	}

	fmt.Println(arrInt)

	slcArrInt := arrInt[0:2]
	fmt.Println(slcArrInt)

	slcArrInt[0] = 3
	fmt.Println(arrInt)

}
