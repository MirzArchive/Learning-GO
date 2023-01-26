package main

import (
	"fmt"
	"strings"
)

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

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for name

	var bookings []string

	for {
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName) // Passing the reference

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		if !isValidName {
			fmt.Println("First Name or Last Name should atleast contain 2 character, try again.")
			continue
		}

		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)

		var isValidEmail bool = strings.Contains(email, "@")
		if !isValidEmail {
			fmt.Println("Email should contain @ character, try again.")
			continue
		}

		fmt.Print("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		var isValidTicketnumber bool = userTickets > 0 && userTickets <= remainingTickets
		if !isValidTicketnumber {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets -= userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		var firstNameSlices []string

		for _, booking := range bookings { // _ -> blank identifier
			names := strings.Fields(booking)
			firstNameSlices = append(firstNameSlices, names[0])
		}
		fmt.Printf("These are all our bookings: %v\n", firstNameSlices)

		if remainingTickets <= 0 {
			fmt.Println("Our conference is booked out. Please come back next year.")
			break
		}

	}

}

func greetUser(name string, quotaTickets uint, availableTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have total of %v tickets and %v are still available\n", quotaTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")
}
