package report

import (
	"airport-basa/Ticket"
	"database/sql"
	"fmt"
)

type AirBase struct {
	db *sql.DB
}

func New(db *sql.DB) AirBase {
	return AirBase{
		db: db,
	}
}

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=davlat password=1 database=airport sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (a AirBase) GetTicketsByCities() ([]ticket.Ticket, error) {
	var From string
	fmt.Println("From: ")
	fmt.Scan(&From)
	var To string
	fmt.Println("To: ")
	fmt.Scan(&To)
	rows, err := a.db.Query(`SELECT id, from_city, to_city, flight_date
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

func GenerateReport() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	airBase := New(db)

	tickets, err := airBase.GetTicketsByCities()
	if err != nil {
		fmt.Println("Error getting tickets:", err)
		return
	}
	fmt.Println("Tickets for the specified cities:")
	for _, t := range tickets {
		fmt.Printf("ID: %d, From: %s, To: %s, Date: %s\n", t.ID, t.FromCity, t.ToCity, t.FlightDate)
	}
}
