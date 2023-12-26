package ticket

import (
	connectdb "airport-basa/connectDB"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Ticket struct {
	ID         int
	FromCity   string
	ToCity     string
	FlightDate string
}
type AirBasa struct {
	db *sql.DB
}

func New(db *sql.DB) AirBasa {
	return AirBasa{
		db: db,
	}
}

func FuncForTickets() {
	db, err :=connectdb.ConnectDB()
	if err != nil {
		log.Fatal("error while connecting to the database")
	}
	defer db.Close()
	airbasa := New(db)
	command := 0
	fmt.Println("Select one of the commands below: ")
	fmt.Print(`
	1 - add ticket
	2 - get ticket by id
	3 - list of tickets
	4 - update ticket by id
	5 - delete ticket by id
	`)
	fmt.Scan(&command)
	switch command {
	case 1:
		if err := airbasa.AddTicket(); err != nil {
			log.Fatal("error AddTicket", err)
		}
	case 2:
		Ticket, err := airbasa.GetTicketByID()
		if err != nil {
			log.Fatal("error during get Ticket by id", err)
		} else {
			fmt.Println(Ticket)
		}
	case 3:
		all, err := airbasa.ListOfTickets()
		if err != nil {
			log.Fatal("error during get all Tickets", err)
		} else {
			fmt.Println(all)
		}
	case 4:
		if err := airbasa.UpdateTicketById(); err != nil {
			log.Fatal("error UpdateTicketById", err)
		}
	case 5:
		if err := airbasa.DeleteTicketByID(); err != nil {
			log.Fatal("Error DeleteTicketByID")
		}
	default:
		fmt.Println("no such command exists...")
	}
}
func (a AirBasa) AddTicket() error {
	var FromCity string = ""
	fmt.Println("from city: ")
	fmt.Scan(&FromCity)
	var ToCity string = ""
	fmt.Println("to city: ")
	fmt.Scan(&ToCity)
	var FlightDate string = ""
	fmt.Println("enter flight date format(02.01.2006): ")
	fmt.Scan(&FlightDate)
	if _, err := a.db.Exec(`insert into ticket (from_city, to_city, flight_date) 
	values ($1, $2, $3)`, FromCity, ToCity, FlightDate); err != nil {
		log.Fatal("error during add Ticket", err)
	}
	fmt.Println("input completed successfully")
	return nil
}
func (a AirBasa) GetTicketByID() (Ticket, error) {
	var TicketId int = 0
	fmt.Println("Enter Ticket id")
	fmt.Scan(&TicketId)
	row := a.db.QueryRow(`select id, from_city, to_city, flight_date from ticket where id = $1`, TicketId)
	GetTicket := Ticket{}
	if err := row.Scan(&GetTicket.ID, &GetTicket.FromCity, &GetTicket.ToCity, &GetTicket.FlightDate); err != nil {
		return Ticket{}, err
	}
	return GetTicket, nil
}

func (a AirBasa) ListOfTickets() ([]Ticket, error) {
	rows, err := a.db.Query(`select id, from_city, to_city, flight_date from ticket`)
	if err != nil {
		return nil, err
	}
	Tickets := []Ticket{}

	for rows.Next() {
		Ticket := Ticket{}
		if err = rows.Scan(&Ticket.ID, &Ticket.FromCity, &Ticket.ToCity, &Ticket.FlightDate); err != nil {
			return nil, err
		}
		Tickets = append(Tickets, Ticket)
	}
	return Tickets, nil
}
func (a AirBasa) UpdateTicketById() error {
	var TicketId int = 0
	fmt.Println("Enter Ticket's id")
	fmt.Scan(&TicketId)
	var TFCity string = ""
	fmt.Println("From city: ")
	fmt.Scan(&TFCity)
	var TTCity string = ""
	fmt.Println("To City: ")
	fmt.Scan(&TTCity)
	var TFlightDate string = ""
	fmt.Println("enter flight date: ")
	fmt.Scan(&TFlightDate)
	if _, err := a.db.Exec(`update ticket set from_city = $1, to_city = $2, flight_date = $3 where id = $4`, TFCity, TTCity, TFlightDate, TicketId); err != nil {
		log.Fatal("error during update Ticket by id", err)
	}
	fmt.Println("update completed successfully")
	return nil
}

func (a AirBasa) DeleteTicketByID() error {
	var TicketId int = 0
	fmt.Println("Enter Ticket id for delete: ")
	fmt.Scan(&TicketId)
	if _, err := a.db.Exec(`delete from ticket where id = $1`, TicketId); err != nil {
		log.Fatal("error during delete ", err)
	}
	fmt.Println("delete completed successfully")
	return nil
}
