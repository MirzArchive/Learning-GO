package main

import "fmt"

func main() {
	var a int = 60
	var b float32
	b = 6.89
	c := "I like eating meat"

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n\n", c, c)

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = conferenceTickets

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, " tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

}
