package main

import "math"

/*
Closure is a function that can be assigned to a variable.
Note: The definition is still ambigous, maybe ill change it later
*/

func main() {
	// Declare basic closure function
	sum := func(a, b int) int {
		return a + b
	}

	// Ways to use a closure function is like any other
	// function, the only difference is to use the variable
	// name instead of the function name (well the function itself doesnt have a name)
	println(sum(1, 2))

	// Immediately-Invoked Function Expression (IIFE)
	// As the name suggest, its a function that immediately invoked/used
	// after it creation like in closure function
	area, circumference := func(d float64) (area, circumference float64) {
		area = math.Pi * math.Pow(d/2, 2)
		circumference = math.Pi * d
		return
	}(9)
	println("Area:", area, "\nCircumference:", circumference)
	println()

	// Closure as return type, basically a function as the return value
	rttAvg, countE2E := func(rtts ...int) (float64, func() float64) {
		var sum int
		for _, rtt := range rtts {
			sum += rtt
		}
		avg := float64(sum / len(rtts))

		return avg, func() float64 {
			return avg / 2
		}
	}(29, 3, 21, 9, 6, 5)
	println(rttAvg)
	println(countE2E())
	println()

	// A function which takes a closure function as the parameter.
	// Generally its called a callback function but im not sure
	pluckEvenNum := func(checker func(num int) bool, numbers ...int) []int {
		evenNum := []int{}
		for _, num := range numbers {
			if checker(num) {
				evenNum = append(evenNum, num)
			}
		}

		return evenNum
	}

	arr := pluckEvenNum(func(num int) bool {
		if num%2 == 0 {
			return true
		}
		return false
	}, 1, 12, 453, 21, 564, 32, 85, 32, 98, 10, 234)

	for _, item := range arr {
		println(item)
	}
}
