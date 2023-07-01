package main

import "math"

func main() {
	printMessage("Halo")
}

// Basic function in Go
func printMessage(msg string) {
	println("Message:", msg)
}

// Function with one or more return type
func sum(a int, b int) (int, bool) {
	sum := a + b

	if sum%2 == 0 {
		return sum, true
	}

	return sum, false
}

// Function parameter with same data type
// can be shortened like this
func greet(greetings, name, emoji string) {
	println(greetings, name, emoji)
}

// Function with predefined return variable
func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2, 2)
	circumference = math.Pi * d

	return
}
