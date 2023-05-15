package main

import "fmt"

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	str := "Hello there"

	for i, s := range str {
		pf("%d : %#U : %c\n", i, s, s)
	}

	pl("Test from Remote Dev Container!")
}
