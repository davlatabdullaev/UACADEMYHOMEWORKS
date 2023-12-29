package report

import (
	connectdb "basa-airways/ConnectDB"
	structforpassengers "basa-airways/StructForPassengers"
	structfortickets "basa-airways/StructForTickets"
	"database/sql"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

type Report struct {
	db *sql.DB
}

func New(db *sql.DB) Report {
	return Report{
		db: db,
	}
}



func (r Report)GetTicketsByCities() {
	db, err := connectdb.Connectdb()
	if err != nil {
		log.Fatal("error while connecting to the database")
	}
	defer db.Close()
	var From, To string
	fmt.Print("From: ")
	fmt.Scan(&From)
	fmt.Print("To: ")
	fmt.Scan(&To)

	rows, err := r.db.Query(`SELECT t.id, t.from_city, t.to_city, p.first_name, p.last_name, p.phone, t.flight_date
		FROM tickets t FULL
		JOIN passengers p ON t.id = p.ticket_id
		WHERE t.from_city = $1 AND t.to_city = $2`, From, To)
	
	defer rows.Close()

	w:= tabwriter.NewWriter(os.Stdout, 1,8,1, '\t', 0)
	fmt.Fprintf(w, "N\tFrom City\tTo City\tPassengers full name \tPhone\tFlight Date\t")
	defer w.Flush()



	for rows.Next() {
		tickets := structfortickets.Tickets{}
		passengers := structforpassengers.Passengers{}
		if err := rows.Scan(&passengers.ID, &tickets.FromCity, &tickets.ToCity, &passengers.FirstName, &passengers.LastName, &passengers.Phone, &tickets.FlightDate); err != nil {
			log.Fatal("error during ... ", err)
		}
		fmt.Fprintf(w, "%d\t%s\t%s\t%s %s\t%s\t%s\n", passengers.ID, tickets.FromCity, tickets.ToCity, passengers.FirstName, passengers.LastName, passengers.Phone, tickets.FlightDate)
	}
	
}
