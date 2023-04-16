package main

import (
	"fmt"
	"strings"
)

func main() {
	conferanceName :="Go Conference"
	const conferanceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application\n",conferanceName)
	fmt.Printf("we have a total of %v tickets and %v are still left.\n",conferanceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for{
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

		isValidName := len(firstname)>=2 && len(lastname)>=2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber :=userTickets>0 && userTickets<=remainingTickets


		if (isValidName && isValidTicketNumber && isValidEmail) {
			remainingTickets = remainingTickets- userTickets
			bookings = append(bookings, firstname + " " +lastname)

			// fmt.Printf("The whole slice : %v\n",bookings)
			// fmt.Printf("The first Value: %v\n",bookings[0])
			// fmt.Printf("Slice type : %T\n",bookings)
			// fmt.Printf("Slice length: %v\n",len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets . You will recieve a confirmation email at %v. \n",firstname,lastname,userTickets,email)
			fmt.Printf("%v tickets remaining for %v \n",remainingTickets,conferanceName)

			firstName := []string{}
			for _,booking := range bookings {
			var names = strings.Fields(booking)
			firstName = append(firstName,names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n",firstName)

			if remainingTickets == 0 {
				fmt.Println("Our all tickets have been booked out")
				break
			
		}
				
		} else{
			if !isValidName {
				fmt.Println("The firstname or lastname is incorrect.Pls correct it Up.")
			}
			if !isValidEmail{
				fmt.Println("The E-mail you entered is not valid . Pls correct it Up.")
			}
			if !isValidTicketNumber{
				fmt.Println("The ticket no You entered up is incorrect , pls correct it up")
			}
			continue
			
	}

	}
}
