package report

import (
	ticket "airport-basa/Ticket"
	connectdb "airport-basa/connectDB"
	"fmt"
	"database/sql"
	"log"
)
type Report struct {
	db *sql.DB
}

func New(db *sql.DB) Report {
	return Report{
		db: db,
	}
}

func (r Report) GetTicketsByCities() ([]ticket.Ticket, error) {
	db, err := connectdb.ConnectDB()
	if err != nil {
		log.Fatal("error while connecting to the database")
	}
	defer db.Close()
	
	var From string
	fmt.Println("From: ")
	fmt.Scan(&From)
	var To string
	fmt.Println("To: ")
	fmt.Scan(&To)
	rows, err := r.db.Query(`SELECT id, from_city, to_city, flight_date
		FROM ticket
		WHERE from_city = $1 AND to_city = $2`, From, To)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := []ticket.Ticket{}
	for rows.Next() {
		ticket := ticket.Ticket{}
		err := rows.Scan(&ticket.ID, &ticket.FromCity, &ticket.ToCity, &ticket.FlightDate)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
	
}

