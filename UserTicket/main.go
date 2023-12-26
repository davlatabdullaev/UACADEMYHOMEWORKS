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
		var password = "12345"
		var code = ""
		fmt.Println("Enter password: ")
		fmt.Scan(&code)
		if code==password{
		passenger.FuncForPassengers()
		} else {
			fmt.Println("wrong password...")
		}
	case 2:
		var parol = "12345"
		var kod = ""
		fmt.Println("Enter password: ")
		fmt.Scan(&kod)
		if kod==parol{
		ticket.FuncForTickets()
		} else {
			fmt.Println("wrong password...")
		}
	case 3:
		db, err := connectdb.ConnectDB()
        if err != nil {
            log.Fatal("Error connecting to the database:", err)
            return
        }
        defer db.Close()
        
        airBaseReport := report.New(db)
        tickets, err := airBaseReport.GetTicketsByCities()
        fmt.Println(tickets)
	default:
		fmt.Println("no such command exists...")
	}
}
