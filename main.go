package main

import "fmt"

func main() {
	hotelName := "MSB hotels and resort"
	const hotelTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Hello, welcome to %v booking app\n", hotelName)
	fmt.Println("Get your hotel tickets here and enjoy your stay")
	fmt.Println("Your comfort is our pleasure.....")

	var firstName string
	var lastName string
	var e_mail string
	var userTickets uint

	// code that prompts you for data:

	fmt.Println("Please enter your firstname:")
	fmt.Scan(&firstName)

	fmt.Println("Your lastname:")
	fmt.Scan(&lastName)

	fmt.Println("Your e-mail address:")
	fmt.Scan(&e_mail)

	fmt.Println("Enter number of Tickets you want:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v Tickets, you'll recieve a confirmation at %v we hope you come back again\n", firstName, lastName, userTickets, e_mail)

	// Ticket reduction logic

	remainingTickets = remainingTickets - userTickets

	// End of ticket reduction logic

	fmt.Printf("%v of %v Tickets left for %v.....!\n", remainingTickets, hotelTickets, hotelName)

}
