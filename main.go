package main

import (
	"GO-Learning/method"
	"fmt"
	"strconv"
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

	method.GreetUser(conferenceName, conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for name

	var bookings = make([]map[string]string, 0)

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

		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		userData["ticket"] = strconv.FormatUint(uint64(userTickets), 10)

		bookings = append(bookings, userData)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		fmt.Printf("These are all our bookings: %v\n", method.GetFirstName(bookings))

		if remainingTickets <= 0 {
			fmt.Println("Our conference is booked out. Please come back next year.")
			break
		}

	}

}
