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

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		firstName, lastName, userEmail, userTickets := getUserInput(remainingTickets)

		isvalidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isvalidName && isValidEmail && isValidTicketNumber {

			remainingTickets, bookings := bookTicket(conferenceName, remainingTickets, userTickets, firstName, lastName, bookings)
			// remainingTickets = remainingTickets - userTickets
			// bookings = append(bookings, firstName+" "+lastName)

			// fmt.Printf("\n User First Name: %v \n User Last Name: %v \n User Email: %v \n booked tickets %v. \n", firstName, lastName, userEmail, userTickets)
			// fmt.Printf("The remaining tickets %v for %v \n", remainingTickets, conferenceName)

			// Call function print first names
			// printFirstNames(bookings)
			firstNames := printFirstNames(bookings)
			fmt.Printf("The first names of the bookings are %v \n", firstNames)
			// var noTicketsRemaining bool = remainingTickets == 0
			// noTicketsRemaining := remainingTickets == 0
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
			// fmt.Println("Invalid Input data")
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets \n", remainingTickets, userTickets)
		}
	}

}

func greetUsers(confName string, confTickets int, remainTickets uint) {
	fmt.Printf("Welcome to %v booking application \n", confName)
	fmt.Printf("We have total tickets of %v & Available are %v \n", confTickets, remainTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings { // can replace index with _ to be unused variables
		var names = strings.Fields(booking) // to slice down string with space
		firstNames = append(firstNames, names[0])
	}

	return firstNames
	// fmt.Printf("The first names of all our bookings: %v \n", firstNames)
	// result := fmt.Sprintf("The first names of all our bookings: %s \n", firstNames)
	// fmt.Printf("The type of entire result is %T \n", result)
}

func validateUserInput(firstName string, lastName string, userEmail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isvalidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isvalidName, isValidEmail, isValidTicketNumber
}

func getUserInput(remainingTickets uint) (string, string, string, uint) {
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

func bookTicket(conferenceName string, remainingTickets uint, userTickets uint, firstName string, lastName string, bookings []string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("\n User First Name: %v \n User Last Name: %v \n booked tickets %v. \n", firstName, lastName, userTickets)
	fmt.Printf("The remaining tickets %v for %v \n", remainingTickets, conferenceName)
	return remainingTickets, bookings
}
