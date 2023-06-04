// Slices for dynamic array
package main

import (
	"fmt"
)

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings []string // creating slice
	bookings := []string{}

	// fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T. \n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total tickets of %v & Available are %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	fmt.Println("Enter number of Tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("\n User First Name: %v \n User Last Name: %v \n User Email: %v \n booked tickets %v. \n", firstName, lastName, userEmail, userTickets)
	// fmt.Printf("Thank you %v %v for booking %v tickets. \nYou will receive a confirmation email at %v \n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("The remaining tickets %v for %v", remainingTickets, conferenceName)

	fmt.Printf("These are all our bookings: %v \n", bookings)
}
