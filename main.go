package main

import (
	"fmt"
	"strconv"
	"strings"
)

var conferanceName = "Go Conference"

const conferanceTickets int = 50

var remainingTickets uint = 50
var bookings = []string{}

type userData struct {
	firstname       string
	lastname        string
	email           string
	numberoOFickets uint
}

func main() {

	greetuser()

	for {
		firstname, lastname, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUser(firstname, lastname, email, userTickets, remainingTickets)

		if isValidName && isValidTicketNumber && isValidEmail {

			bookings := bookTicket(remainingTickets, userTickets, bookings, firstname, lastname, email, conferanceName)
			FirstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", FirstNames)

			firstName := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstName = append(firstName, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstName)

			if remainingTickets == 0 {
				fmt.Println("Our all tickets have been booked out")
				break

			}

		} else {
			if !isValidName {
				fmt.Println("The firstname or lastname is incorrect.Pls correct it Up.")
			}
			if !isValidEmail {
				fmt.Println("The E-mail you entered is not valid . Pls correct it Up.")
			}
			if !isValidTicketNumber {
				fmt.Println("The ticket no You entered up is incorrect , pls correct it up")
			}
			continue

		}

	}
}

func greetuser() {
	fmt.Printf("Welcome to %v booking application\n", conferanceName)
	fmt.Printf("we have a total of %v tickets and %v are still left.\n", conferanceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string {
	firstName := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstName = append(firstName, names[0])
	}
	return firstName

}
func validateUser(firstname string, lastname string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
func getUserInput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastname)

	fmt.Println("Enter your E-mail: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of Tickets: ")
	fmt.Scan(&userTickets)
	return firstname, lastname, email, userTickets
}

func bookTicket(remainingTickets uint, userTickets uint, bookings []string, firstname string, lastname string, email string, conferanceName string) []string {
	remainingTickets = remainingTickets - userTickets

	//created a map for the user
	var userdata = make(map[string]string)
	userdata["firstName"] = firstname
	userdata["lastName"] = lastname
	userdata["email"] = email
	userdata["numberofTickets"] = strconv.FormatUint(uint64(userTickets),10)

	bookings = append(bookings, firstname+" "+lastname)
	fmt.Printf("Thank you %v %v for booking %v tickets . You will recieve a confirmation email at %v. \n", firstname, lastname, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferanceName)
	return bookings
}
