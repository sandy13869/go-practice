package main

import (
	"fmt"
	"strings"
)

func greetUsers() {
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total tickets of %v & Available are %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint
	// ask for their name
	fmt.Println("Enter your First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email: ")
	fmt.Scan(&userEmail)

	fmt.Printf("Available Tickets %v \n", remainingTickets)
	fmt.Println("Enter number of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func validateUserInput(firstName string, lastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isvalidName, isValidEmail, isValidTicketNumber
}

func printFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings { // can replace index with _ to be unused variables
		var names = strings.Fields(booking) // to slice down string with space
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, bookings []string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("\n User First Name: %v \n User Last Name: %v \n booked tickets %v. \n", firstName, lastName, userTickets)
	fmt.Printf("The remaining tickets %v for %v \n", remainingTickets, conferenceName)
	return remainingTickets, bookings
}
