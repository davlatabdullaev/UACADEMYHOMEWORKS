package main

import (
	passenger "airport-basa/Passenger"
	report "airport-basa/Report"
	ticket "airport-basa/Ticket"
	connectdb "airport-basa/connectDB"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
		db, err := connectdb.ConnectDB()
		if err != nil {
			log.Fatal("Error connecting to the database:", err)
			return
		}
		defer db.Close()
	
		airBaseReport := report.New(db)
		if err := airBaseReport.GetTicketsByCities(); err != nil {
			log.Fatal("Error generating report:", err)
		}
}

}
