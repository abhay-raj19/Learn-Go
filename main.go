package main

import "fmt"

func main() {
	conferanceName :="Go Conference"
	const conferanceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Printf("Welcome to %v booking application\n",conferanceName)
	fmt.Printf("we have a total of %v tickets and %v are still left.\n",conferanceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")

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

	remainingTickets = remainingTickets- userTickets
	bookings[0] = firstname + " " + lastname

	fmt.Printf("The whole array : %v\n",bookings)
	fmt.Printf("The first Value: %v\n",bookings[0])
	fmt.Printf("Array type : %T\n",bookings)
	fmt.Printf("Array length: %v\n",len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets . You will recieve a confirmation email at %v. \n",firstname,lastname,userTickets,email)
	fmt.Printf("%v tickets remaining for %v \n",remainingTickets,conferanceName)

}
