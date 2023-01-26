package main

import (
	"fmt"
	"strings"
)

func greetUser(name string, quotaTickets uint, availableTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have total of %v tickets and %v are still available\n", quotaTickets, availableTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName(bookingNames []string) []string {
	var firstName []string
	for _, booking := range bookingNames { // _ -> blank identifier
		name := strings.Fields(booking)
		firstName = append(firstName, name[0])
	}
	return firstName
}