package main

import "fmt"

func main() {
	/*var a int = 60
	var b float32
	b = 6.89
	c := "I like eating meat"

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n\n", c, c)*/

	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string
	var userTickets int
	// ask user for name

	fmt.Println(&conferenceName)

	fmt.Scan(userName)
	userTickets = 3
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
