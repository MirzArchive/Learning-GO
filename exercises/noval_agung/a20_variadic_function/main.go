package main

/*
Variadic function is just like a normal function but
it can have unlimited arguments in its parameter
*/

func main() {
	// Example usage of an variadic function
	println(sum(1, 2))
	println(sum(1, 2, 3, 4, 5, 6, 7, 8))
	println(sum(1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15))
	println(sum(1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))

	// Variadic Function can also take a single slice
	// as its argument instead of multiple argument
	numbers := []int{1, 2, 3, 4, 5, 6}
	println(sum(numbers...))
}

// Basic variadic function to sum all passed number
func sum(numbers ...int) int {
	var total int
	for _, num := range numbers {
		total += num
	}
	return total
}
