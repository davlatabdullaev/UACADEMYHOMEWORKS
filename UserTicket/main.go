package main

import (
	passenger "airport-basa/Passenger"
	report "airport-basa/Report"
	ticket "airport-basa/Ticket"
	"fmt"
	_"github.com/lib/pq"
)

func main() {
	cmd := 0
	fmt.Println("Select the list")
	fmt.Print(`
	1 - User
	2 - Ticket
	3 - Report
	`)
	fmt.Scan(&cmd)
	switch cmd {
	case 1:
		passenger.FuncForPassengers()
	case 2:
		ticket.FuncForTickets()
	case 3:
		report.GenerateReport()
	default:
		fmt.Println("no such command exists...")
	}
}
