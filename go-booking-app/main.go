package main

import (
	"bufio"
	"fmt"
	"go-booking-app/helper"
	"os"
	"strconv"
	"strings"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

/*
*  Accessing the values in global scope like - public static final variable
while declaring here, we cannot use type inference :=

Best practice - always declare variable as local as neede
*/
var conferenceName = "Go Conference"

const noOfTickets int = 50       // total tickets unchangable so kept as constant
var remainingTickets uint32 = 50 // values are changeable so kept in var
var bookings []string
var mapOfBookings = make([]map[string]string, 0) // List of Maps => ParentKey : [{key : value }, { key : value}, {key : value}]
// initialized list of maps size as Zero (0)

// use struct like model class in map data
var structUserDetail = make([]UserDetail, 0) // List of Maps => ParentKey : [{key : value }, { key : value}, {key : value}]

type UserDetail struct {
	userName    string
	email       string
	mobile      string
	userTickets int
	isEnabled   bool
}

// wait group is like Completable Future, wait of completion
//var wg = sync.WaitGroup{} // this gives synchronization

func main() {

	/**
	Variables Example:

	1) var remainingTicketsUint uint16 = -50 // if we specify uint then it cannot be from negative so its
	 starts from 0 to 65535 but if you kept as int then it can be from -65535 to 65535

	2) var conferenceName string = "Go Conference" // we can specify datatype explicit instead of go detecting it

	3)Syntactic sugar : only for variables without explicit data type we can use below syntax, it cannot be done in const

		conferenceNameSyntaticSugar := "Go Conference"
	*/

	/** Data types

	String, Integer (int8, int16, int32, int64 vs uint8, ...), float (float32, float64 only), complex(complex64, complex128),
	Maps, Boolean, Arrays

	*/
	//conferenceName := "Go Conference"
	//const noOfTickets int = 50       // total tickets unchangable so kept as constant
	//var remainingTickets uint32 = 50 // values are changeable so kept in var

	var userName string // need to initialize data type when variable is created
	var userTickets uint32
	var email string
	var mobile string

	/**
	Arrays / Slices

	var bookings [50]string => array

	var bookings []string => slice dynamic in size and its types of implementations
	bookings := []string{}
	var bookings = []string{"chandru"}



	Examples :
	var bookings [50]string - empty declaration without initializing any values
	var bookings = [50]string{"chandru", "gowtham", "jeeva"}
	var bookingsEmptyArr = [50]string{}

	-- To add value to array
	bookings[0] = "value1"

	-- To add value to slice
	bookings = append(bookings, "value1")

	-- useful functions
	fmt.Printf("Booking array %v", bookings[0])
	fmt.Printf("Booking array %T", bookings)
	fmt.Printf("Booking array length %v", len(bookings))

	*/

	/**
	  Print Statements Example:

	1) while Printf use %v everywhere , %s prints object as string, %T prints datatype of variable
	  2) Printf is used to print values dynamically, formatted data
	  3) Println is used to print constant sentences with line space or use like below args

	fmt.Printf("**** Hi Everyone, Welcome to our %v ****\n\n", conferenceName)
	fmt.Printf("**** Hi Everyone, Welcome to our %s ****\n\n", conferenceName)
	fmt.Printf("Hello %v congrats on bookings %v tickets for %v \n", userName, userTickets, conferenceName)
	fmt.Printf("Hello %T congrats on bookings %T tickets for %T \n", userName, userTickets, conferenceName)

	*/

	/**
	Loops:
	it has only one for loop in go-lang, but its has different types of implementations to satisfy all the needs

	*/

	for {

		helper.GreetUsers(conferenceName, noOfTickets, remainingTickets)

		fmt.Println("Enter your Name!")
		reader := bufio.NewReader(os.Stdin)
		userName, _ = reader.ReadString('\n')

		fmt.Println("Enter your Email!")
		fmt.Scan(&email)

		fmt.Println("Enter your MobileNumber!")
		fmt.Scan(&mobile)

		fmt.Println("Enter how many tickets needed for", conferenceName)
		fmt.Scan(&userTickets)

		/**
		Pointers - special variables

		1) stored memory address

		//fmt.Println(&userName) // prints memory-address for the variable called as pointers
		//fmt.Println(&userTickets) // prints memory-address for the variable as pointers

		*/

		/**
		if else conditions

		similar like java

		for userTickets > remainingTickets  && userTickets != 0  && remainingTickets != 0 {
		}

		for true { // infinite loop
		}
		*/

		/**
		Switch statements
		*/

		switch userTickets {

		case 1:
			fmt.Println("You have booking only 1 ticket , book more to get offer!")

		case 5, 10:
			fmt.Println("You have booking only 5 ticket , book more to get offer!")
		default:
			fmt.Println("You have not applied any offer!")

		}

		if userTickets > remainingTickets {
			fmt.Printf("Conference doesn't have these count of tickets %v", userTickets-remainingTickets)
			//break  // exit whole for loop
			continue // just skip one iteration by not executing next lines
		} else {
			fmt.Println("Conference has enough tickets")
		}

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, strings.TrimSuffix(userName, "\n"))

		/**
		Maps
		userName as key
		username, email, mobile, tickest as list of values
		//create a map and initialize
		var userData map[string]string
		var userData  = map[string]string

		var mapBookings = make(map[string]string) // Creating empty map

		*/

		var userData = make(map[string]string) // Creating empty map
		userData["name"] = userName
		userData["email"] = email
		userData["mobile"] = mobile
		userData["tickets"] = strconv.FormatUint(uint64(int(userTickets)), 10)

		mapOfBookings = append(mapOfBookings, userData)

		var userDataStruct = UserDetail{
			userName:    userName,
			email:       email,
			mobile:      mobile,
			userTickets: int(userTickets),
			isEnabled:   true,
		}

		structUserDetail = append(structUserDetail, userDataStruct)

		fmt.Printf("User data List of Struct: %v\n", structUserDetail)

		fmt.Printf("User data inside single object of Struct: %v\n", structUserDetail[0].userTickets)

		/**
		Loops: Get all names of user in separate list from bookings

		//for _, name := range bookings { // if variables are not used use Underscore as their keyword
		//	var names = strings.Fields(name) // by default this split string text with space as character and return a Slice
		//	firstNames = append(firstNames, names[0])
		//}

		*/

		newBookings, _, _ := helper.PrintInLoop(bookings)
		fmt.Printf("Bookings whole data: %v\n", newBookings)

		fmt.Printf("Map of Bookings whole data: %v\n", mapOfBookings)
		_ = helper.PrintMapLoop(mapOfBookings)

		fmt.Printf("Congrats %v on bookings %v tickets for conference ! Available tickets => %v\n", strings.TrimSuffix(userName, "\n"), userTickets, remainingTickets)

		var stillConferenceAvailable bool = remainingTickets == 0
		fmt.Printf("Conference is house full True / False: %v", stillConferenceAvailable)

		//wg.Add(1) // add threads to wait
		go helper.SendTicket(userName)
		//wg.Wait() // ask thread to wait

		if stillConferenceAvailable {
			fmt.Printf("Conference is house full")
			break
		}

	}

}
