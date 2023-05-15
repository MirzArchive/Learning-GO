package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pL = fmt.Println

func main() {
	rd := bufio.NewReader(os.Stdin)

	fmt.Print("Whats your name? ")
	name, err := rd.ReadString('\n')
	if err != nil {
		log.Fatalln(err.Error())
	}

	pL("Hello", name)
}
