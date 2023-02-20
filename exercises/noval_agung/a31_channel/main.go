package main

import "fmt"

/*
A channel is a tunnel to connect goroutine with other goroutine.
This tunnel are used to transfer or receive data. The process of transfering
and receiving with channel are blocking the entire program.
*/

func main() {
	// Basic usage of channel
	msgChan := make(chan string)

	greet := func(name string) {
		message := fmt.Sprintf("Hello %s", name)
		msgChan <- message
	}

	go greet("Mirza")
	go greet("Ghatfan")
	go greet("Fajru")

	msg1 := <-msgChan
	fmt.Println(msg1)

	msg2 := <-msgChan
	fmt.Println(msg2)

	msg3 := <-msgChan
	fmt.Println(msg3)
}
