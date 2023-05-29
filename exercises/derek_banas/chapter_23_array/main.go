package main

import "fmt"

func main() {
	arrInt := [2][2]int{{1, 3}, {2, 6}}

	for _, a := range arrInt {
		for _, b := range a {
			fmt.Println(b)
		}
	}
}
