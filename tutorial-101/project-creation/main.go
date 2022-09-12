package main

import "fmt"

func main() {

	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to", conferenceName, "booking Application")
	fmt.Println("Get your tickets attended here!! We have a total of ", conferenceTickets, "tickets!!")
	fmt.Println("We have remaining", remainingTickets, "tickets")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//getting the username from user
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets which you want to book:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v!! You have requested for %v tickets, you will receive confirmation at %v\n", firstName, lastName, userTickets, email)
}
