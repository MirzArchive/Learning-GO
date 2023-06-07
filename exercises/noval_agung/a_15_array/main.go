package main

import "fmt"

func main() {
	var arr1 [3]int
	arr1 = [3]int{1, 2, 3}

	var _ = [3]int{
		4,
		5,
		6,
	}

	// NOT array, but a slice
	var _ = make([]int, 3)
	_ = []int{7, 8, 9}

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	arr3 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	arr4 := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}

	for i := range arr3 {
		fmt.Printf("arr 3: %v, arr4: %v\n", arr3[i], arr4[i])
	}
}
