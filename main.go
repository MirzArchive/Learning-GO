package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func sum[T constraints.Ordered](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(1.1, 2.2))
}
