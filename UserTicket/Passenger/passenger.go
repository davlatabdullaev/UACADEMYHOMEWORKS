package passenger

import (
	connectdb "airport-basa/connectDB"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Passenger struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Phone     string
	TicketID  int
}
type AirBasa struct {
	db *sql.DB
}

func New(db *sql.DB) AirBasa {
	return AirBasa{
		db: db,
	}
}

func FuncForPassengers() {
	db, err :=connectdb.ConnectDB()
	if err != nil {
		log.Fatal("error while connecting to the database")
	}
	defer db.Close()
	airbasa := New(db)
	command := 0
	fmt.Println("Select one of the commands below: ")
	fmt.Print(`
	1 - add passenger
	2 - get passenger by id
	3 - list of passengers
	4 - update passenger by id
	5 - delete passenger by id
	`)
	fmt.Scan(&command)
	switch command {
	case 1:
		if err := airbasa.AddPassenger(); err != nil {
			log.Fatal("error AddPassenger", err)
		}
	case 2:
		passenger, err := airbasa.GetPassengerByID()
		if err != nil {
			log.Fatal("error during get passenger by id", err)
		} else {
			fmt.Println(passenger)
		}
	case 3:
		all, err := airbasa.ListOfPassengers()
		if err != nil {
			log.Fatal("error during get all passengers", err)
		} else {
			fmt.Println(all)
		}
	case 4:
		if err := airbasa.UpdatePassengerByID(); err != nil {
			log.Fatal("error UpdatePassengerByID", err)
		}
	case 5:
		if err := airbasa.DeletePassengerByID(); err != nil {
			log.Fatal("Error DeletePassengerByID")
		}
	default:
		fmt.Println("no such command exists...")
	}
}
func (a AirBasa) AddPassenger() error {
	var FirstName string = ""
	fmt.Println("enter passenger's first name: ")
	fmt.Scan(&FirstName)
	var LastName string = ""
	fmt.Println("enter passenger's last name: ")
	fmt.Scan(&LastName)
	var Email string = ""
	fmt.Println("enter passenger's email: ")
	fmt.Scan(&Email)
	var Phone string = ""
	fmt.Println("enter passenger's phone number: ")
	fmt.Scan(&Phone)
	var TicketID int = 0
	fmt.Println("enter passenger's ticket id: ")
	fmt.Scan(&TicketID)

	if _, err := a.db.Exec(`insert into passenger (first_name, last_name, email, phone, ticket_id) 
	values ($1, $2, $3, $4, $5)`, FirstName, LastName, Email, Phone, TicketID); err != nil {
		log.Fatal("error during add passenger", err)
	}
	fmt.Println("input completed successfully")
	return nil
}
func (a AirBasa) GetPassengerByID() (Passenger, error) {
	var PassengerId int = 0
	fmt.Println("Enter passenger id")
	fmt.Scan(&PassengerId)
	row := a.db.QueryRow(`select id, first_name, last_name, email, phone, ticket_id from passenger where id = $1`, PassengerId)
	GetPassenger := Passenger{}
	if err := row.Scan(&GetPassenger.ID, &GetPassenger.FirstName, &GetPassenger.LastName, &GetPassenger.Email, &GetPassenger.Phone, &GetPassenger.TicketID); err != nil {
		return Passenger{}, err
	}
	return GetPassenger, nil
}

func (a AirBasa) ListOfPassengers() ([]Passenger, error) {
	rows, err := a.db.Query(`select * from passenger`)
	if err != nil {
		return nil, err
	}
	Passengers := []Passenger{}

	for rows.Next() {
		passenger := Passenger{}
		if err = rows.Scan(&passenger.ID, &passenger.FirstName, &passenger.LastName, &passenger.Email, &passenger.Phone, &passenger.TicketID); err != nil {
			return nil, err
		}
		Passengers = append(Passengers, passenger)
	}
	return Passengers, nil
}
func (a AirBasa) UpdatePassengerByID() error {
	var PassengerId int = 0
	fmt.Println("Enter passenger id")
	fmt.Scan(&PassengerId)
	var PFName string = ""
	fmt.Println("Enter new passenger's first name ")
	fmt.Scan(&PFName)
	var PLName string = ""
	fmt.Println("Enter new passenger's last name ")
	fmt.Scan(&PLName)
	var PEmail string = ""
	fmt.Println("Enter new passenger's email ")
	fmt.Scan(&PEmail)
	var PPhone string = ""
	fmt.Println("Enter new passenger's phone number")
	fmt.Scan(&PPhone)
	if _, err := a.db.Exec(`update passenger set first_name = $1, last_name = $2, email = $3, phone = $4 
	where id = $5`, PFName, PLName, PEmail, PPhone, PassengerId); err != nil {
		log.Fatal("error during update passenger by id", err)
	}
	fmt.Println("update completed successfully")
	return nil
}

func (a AirBasa) DeletePassengerByID() error {
	var PassengerId int = 0
	fmt.Println("Enter passenger id for delete: ")
	fmt.Scan(&PassengerId)
	if _, err := a.db.Exec(`delete from passenger where id = $1`, PassengerId); err != nil {
		log.Fatal("error during delete ", err)
	}
	fmt.Println("delete completed successfully")
	return nil
}
