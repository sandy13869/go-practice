// Mapping data
package main

import (
	"fmt"
	"strconv"
)

// var conferenceTickets = os.Getenv("AVAILABLE_TICKETS")
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) // define a mapping component with initialization

// var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, userEmail, userTickets := getUserInput()

		isvalidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

		if isvalidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, userEmail)

			firstNames := printFirstNames()
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
				fmt.Println("Invalid Email Address (missing @).")
			}
			if !isValidTicketNumber {
				fmt.Println("Entered Invalid Ticket Number.")
			}
		}
	}

}

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

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // can replace index with _ to be unused variables
		// to slice down string with space
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for the user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["userMail"] = userEmail
	userData["ticketsBooked "] = strconv.FormatUint(uint64(userTickets), 10) // Implicitly convert uint into the string for mapping
	// userData["userTickets"] = userTickets

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v \n", bookings)

	fmt.Printf("\n User First Name: %v \n User Last Name: %v \n booked tickets %v. \n", firstName, lastName, userTickets)
	fmt.Printf("The remaining tickets %v for %v \n", remainingTickets, conferenceName)

}
