package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Gaysang's Bootcamp"
	const conferenceTicket uint = 50
	var remainingTicket uint = conferenceTicket
	bookings := []string{}

	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTicket, remainingTicket)
	fmt.Printf("Get Your Ticket\n")

	for {
		var firstName string
		var lastName string
		var email string
		var userTicket int

		fmt.Println("Enter Your First Name")
		fmt.Scan(&firstName)

		fmt.Println("Enter Your Last Name")
		fmt.Scan(&lastName)

		fmt.Println("Enter Your Email Address")
		fmt.Scan(&email)

		fmt.Println("Enter Number of Tickets:")
		fmt.Scan(&userTicket)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTicket := userTicket > 0 && uint(userTicket) <= remainingTicket

		if isValidEmail && isValidName && isValidTicket {
			remainingTicket -= uint(userTicket)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank You %v %v for Booking %v Tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
			fmt.Printf("%v Tickets Remaining at %v Booking Application\n", remainingTicket, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These Are All Our Bookings: %v\n", firstNames)

			noTicketRemaining := remainingTicket == 0
			if noTicketRemaining {
				fmt.Printf("%v is Sold Out. Come back next century.\n", conferenceName)
				break
			}

		} else {
			if !isValidName {
				fmt.Println("Your name is too short")
			}

			if !isValidEmail {
				fmt.Println("Invalid email address")
			}

			if !isValidTicket {
				fmt.Println("Invalid Number of Tickets")
			}
		}
	}

	// Ini contoh switch case
	fmt.Println("Enter your Country:")
	var city string
	fmt.Scan(&city)

	switch city {
	case "Indonesia":
		fmt.Println("Your Country is a Indonesia")
	case "Malaysia":
		fmt.Println("Your Country is a Malaysia")
	case "Thailand":
		fmt.Println("Your Country is a Thailand")
	case "Israel":
		fmt.Println("F off, GO TO HELL!!!")
	default:
		fmt.Println("No valid city selected")
	}

}
