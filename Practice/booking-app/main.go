// Struct data mixed datatypes and concurrency
package main

import (
	"fmt"
	"sync"
	"time"
)

// var conferenceTickets = os.Getenv("AVAILABLE_TICKETS")
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // define a Struct with initialization

type UserData struct {
	firstName     string
	lastName      string
	userEmail     string
	ticketsBooked uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for {

	firstName, lastName, userEmail, userTickets := getUserInput()

	isvalidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets)

	if isvalidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, userEmail)

		wg.Add(1)                                                  // wait the thread for execute tail function
		go sendTicket(userTickets, firstName, lastName, userEmail) // to make our application concurent

		firstNames := printFirstNames()
		fmt.Printf("The first names of the bookings are %v \n", firstNames)

		if remainingTickets == 0 {
			// end programm
			fmt.Println("Our conference is booked out, Try again")
			// break
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
	// }
	wg.Wait() // Blocks until the Wait Group counter is 0

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
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets

	// create a struct for the userdata
	var userData = UserData{
		firstName:     firstName,
		lastName:      lastName,
		userEmail:     userEmail,
		ticketsBooked: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of Bookings is %v \n", bookings)

	fmt.Printf("\n User First Name: %v \n User Last Name: %v \n booked tickets %v. \n", firstName, lastName, userTickets)
	fmt.Printf("The remaining tickets %v for %v \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("#---------------------#")
	fmt.Printf("Sending ticket %v to email address %v \n", ticket, userEmail)
	fmt.Println("#---------------------#")
	wg.Done() // removes the Wait Group block and reset
}
