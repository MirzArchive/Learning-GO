package main

import "fmt"

func main() {
	// Slice is a reference data type of one or more array element

	// Slice declaration and initialization
	// -------------------------------
	// var slice1 []int
	// slice1 = []int{1, 2, 3}
	// -------------------------------
	// slice2 := make([]string, 3)
	// -------------------------------
	// slice3 := []string{"durian", "banana", "lemon", "grape"}
	// -------------------------------
	// slice4 := arr[0:4]
	// -------------------------------

	arr := [...]string{"watermelon", "pinnaple", "apple", "orange"}
	slice6 := arr[0:4]

	// Proof that Slice is an reference of array elements

	fmt.Println(slice6) // [watermelon pinnaple apple orange]

	arr[0] = "dragon_fruit"

	fmt.Println(slice6) // [dragon_fruit pinnaple apple orange]

	// Difference of len() and cap()
	fromZero := arr[0:2]
	fmt.Println("len:", len(fromZero), "| cap:", cap(fromZero)) // len: 2 | cap: 4 -> anything start from 0 always take the original array length as cap

	fromOne := arr[1:3]
	fmt.Println("len:", len(fromOne), "| cap:", cap(fromOne)) // len: 2 | cap: 3 -> i don't even know how's this 3
	fmt.Println(fromOne[1])

	// Appending new value into Slices with append()
	mrSlice := arr[0:4]
	fmt.Println(mrSlice) // [dragon_fruit pinnaple apple orange]

	mrSlice = append(mrSlice, "buddha's_hand", "kiwi")
	fmt.Println(mrSlice) // [dragon_fruit pinnaple apple orange buddha's_hand kiwi]

	// About copy()
	dst := make([]string, 2)
	src := arr[1:4]

	copy(dst, src)

	fmt.Println("dst:", dst, "| src:", src) // dst: [pinnaple apple] | src: [pinnaple apple orange]

	// Accessing Slice or Array Element with 3 Index. By using this it also specify the cap
	withNoCap := arr[0:2]
	withCap := arr[0:2:2]

	fmt.Println("len:", len(withNoCap), "| cap:", cap(withNoCap)) // len: 2 | cap: 4

	fmt.Println("len:", len(withCap), "| cap:", cap(withCap)) // len: 2 | cap: 2
}
