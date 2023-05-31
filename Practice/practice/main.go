// Packages
package main

import (
	"fmt"
	"os"
)

var conferenceTickets = os.Getenv("AVAILABLE_TICKETS")

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings []string

func main() {

	greetUsers()

	for {

		firstName, lastName, userEmail, userTickets := getUserInput()

		isvalidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userEmail, userTickets)

		if isvalidName && isValidEmail && isValidTicketNumber {

			remainingTickets, bookings := bookTicket(userTickets, firstName, lastName, bookings)

			firstNames := printFirstNames(bookings)
			fmt.Printf("The first names of the bookings are %v \n", firstNames)

			if remainingTickets == 0 {
				// end programm
				fmt.Println("Our conference is booked out, Try again")
				break
			}
			// } else if {
		} else {
			if !isvalidName {
				fmt.Println("First Name or Last Name u entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email Address.")
			}
			if !isValidTicketNumber {
				fmt.Println("Entered Invalid Ticket Number.")
			}
		}
	}

}
