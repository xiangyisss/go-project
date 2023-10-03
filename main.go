package main

import (
	"Go-lang-project/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50
var conferenceName = "gophercon"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	// for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail,isValidUserTicketsNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail &&  isValidUserTicketsNumber {
			
			bookTickets(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTickets(userTickets, email, firstName, lastName)

			firstNames := getFristNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out! Coma back next year.")
				// break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			
			if !isValidEmail {
				fmt.Println("Email address you entered doens't contain a @ sign.")
			}

			if !isValidUserTicketsNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	// }
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFristNames () []string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames =  append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput () (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	
	fmt.Println("Enter number of tickets you want to book: ")
	fmt.Scan(&userTickets)
	
	return firstName, lastName, email, userTickets
}

func bookTickets (userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	
	var userData = userData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of booking is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTickets (userTickets uint, email string, firstName string, lastName string) {
	time.Sleep(30 * time.Second)
	fmt.Println("#################")
	var ticket = fmt.Sprintf("%v Tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending tickets: \n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}