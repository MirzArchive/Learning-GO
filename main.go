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

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for name

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName) // Passing the reference

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
}
