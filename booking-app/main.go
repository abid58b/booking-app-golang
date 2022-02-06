package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
	"sync"
	"time"
)

var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name

	fmt.Print("Enter Your First Name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter Your Last Name: ")
	fmt.Scan(&lastName)

	isValidName := len(firstName) > 2 && len(lastName) > 2

	if !isValidName {
		fmt.Println("You Enter first name or Last Name is too Short Try Again")

	} else {
		fmt.Print("Enter Your Email Address: ")
		fmt.Scan(&email)

		isValidEmail := strings.Contains(email, "@")

		if !isValidEmail {
			fmt.Println("You Enter Wrong Email")

		} else {
			fmt.Print("How many Tickets do you Want: ")
			fmt.Scan(&userTickets)

			isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

			if !isValidTicketNumber {
				fmt.Println("Number of Ticket you entered is invalid")
			} else {
				if isValidName && isValidEmail && isValidTicketNumber {

					bookTickets(remainingTickets, userTickets, firstName, lastName, email, conferenceName)
					wg.Add(1)
					go sendTicket(userTickets, firstName, lastName, email)

					fmt.Printf("The First Name of Bookings are: %v\n", getFirstNames())

					if remainingTickets == 0 {
						//end program
						fmt.Printf("%v is booked out. Come Back Next Year.\n", conferenceName)

					}

				} else {
					fmt.Printf("We have only %v tickets remaining. So we can't book %v tickets.\n", remainingTickets, userTickets)

				}

			}
		}
	}

	wg.Wait()
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames

}
func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of Booking Data is %v \n", bookings)

	fmt.Printf("Thank you %v %v for Register %v Tickets. You will Recieved confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket: \n %v \n to email address %v \n", ticket, email)
	fmt.Println("###############")
	wg.Done()
}
