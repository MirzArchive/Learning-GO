package main

import (
	"fmt"
	"runtime"
)

/*
By default, every transfer operation on channel are unbuffered.
What it really means that,any other statement after channel data
transfer statement would not be execute. By using a buffered channel,
the process to transfer a data through channel are asynchronous while
the process of receiving or the receiver ar always synchronous.
*/

func main() {
	runtime.GOMAXPROCS(2)

	messages := make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
