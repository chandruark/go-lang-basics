package helper

import (
	"fmt"
	"strings"
	"time"
)

//var ConferenceNameGlobal = "Hello go " //global scope because it starts with Caps

func GreetUsers(conferenceName string, noOfTickets int, remainingTickets uint32) {
	fmt.Printf("\n\n**** Hi Everyone, Welcome to our %s ****\n\n", conferenceName)
	fmt.Println("No of Tickets for", conferenceName, "is", noOfTickets, "and Remaining tickets", remainingTickets)
	fmt.Println("Please Enter your personal details to book tickets!")
}

func PrintInLoop(bookings []string) ([]string, string, string) {

	firstNames := []string{}

	for _, name := range bookings { // if variables are not used use Underscore as their keyword
		var names = strings.Fields(name) // by default this split string text with space as character and return a Slice
		firstNames = append(firstNames, names[0])
	}
	return firstNames, firstNames[0], firstNames[0]
}

func PrintMapLoop(mapOfBookings []map[string]string) []string {

	firstNames := []string{}

	for _, mapValues := range mapOfBookings { // if variables are not used use Underscore as their keyword
		fmt.Println("Map User data", mapValues["name"], mapValues["email"], mapValues["tickets"])

	}
	return firstNames
}

func SendTicket(userName string) {
	time.Sleep(20 * time.Second)
	var template = fmt.Sprintf("send an email %v", userName)
	fmt.Println(template)

}
