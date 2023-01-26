package method

import (
	"fmt"
)

func GreetUser(name string, quotaTickets uint, availableTickets uint) { // -> Pascal Case naming for exporting
	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have total of %v tickets and %v are still available\n", quotaTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")
}

func GetFirstName(bookings map[string]string) []string {
	var firstName []string
	for key, booking := range bookings { // _ -> blank identifier
		switch key {
		case "firstName":
			firstName = append(firstName, booking)
		default:
			continue
		}
	}
	return firstName
}