package helper

import "fmt"

func GreetUsers(confName string, confTickets int, remainTickets uint) {

	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total %v tickets and %v are still available\n", confTickets, remainTickets)
	fmt.Println("Get your Ticket  here to attend")

}
