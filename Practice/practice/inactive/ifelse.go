// Slices for dynamic array
package main

import (
	"fmt"
	"strings"
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

	for {
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

		// if userTickets > remainingTickets {
		// 	fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
		// 	// break // ite breaks entire program
		// 	continue // it will takes control to the start without break the program
		// }

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\n User First Name: %v \n User Last Name: %v \n User Email: %v \n booked tickets %v. \n", firstName, lastName, userEmail, userTickets)
			// fmt.Printf("Thank you %v %v for booking %v tickets. \nYou will receive a confirmation email at %v \n", firstName, lastName, userTickets, userEmail)
			fmt.Printf("The remaining tickets %v for %v \n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { // can replace index with _ to be unused variables
				var names = strings.Fields(booking) // to split string with space
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of all our bookings: %v \n", firstNames)

			// var noTicketsRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0
			if remainingTickets == 0 {
				// end programm
				fmt.Println("Our conference is booked out, Try again")
				break
			}
			// } else if {
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
		}
	}

}
