package commandlist

import (
	funcsforpassengers "basa-airways/FuncsForPassengers"
	funcsfortickets "basa-airways/FuncsForTickets"
	report "basa-airways/Report"
	"fmt"
)

func CommandList(){
	command := 0
	fmt.Println("Select command: ")
	fmt.Print(`
	1 - Passengers
	2 - Tickets
        3 - Report
	`)
	fmt.Scan(&command)
	switch(command) {
	case 1: 
	// Passengers
	funcsforpassengers.PassengersFunc()
	case 2: 
	// Tickets
	funcsfortickets.FuncsForTickets()
	case 3:
	// Reports
	report.Report{}.GetTicketsByCities()
	default:
		fmt.Println("no such command exists...")	
	}

}