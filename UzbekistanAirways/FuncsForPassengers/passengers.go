package funcsforpassengers

import (
	connectdb "basa-airways/ConnectDB"
	structforpassengers "basa-airways/StructForPassengers"
	"database/sql"
	"fmt"
	"log"
)

type BasaPassengers struct {
	db *sql.DB
}

func New(db *sql.DB) BasaPassengers {
	return BasaPassengers{
		db: db,
	}
}

func PassengersFunc() {
	db, err := connectdb.Connectdb()
	if err != nil {
		log.Fatal("error during connect database...",err)
	}
	defer db.Close()
	command := 0
	BasaPassenger := New(db)
	fmt.Println("Select command: ")
	fmt.Print(`
	1 - Insert Passenger
	2 - Get Passenger By Id
	3 - List of Passengers
	4 - Update Passenger By Id
	5 - Delete Passenger By ID
	`)
	fmt.Scan(&command)
	switch(command) {
	case 1:
		// Insert Passenger
		if err := BasaPassenger.InsertPassenger(); err != nil {
			log.Fatal("error while insert passenger... ", err)
		}
		fmt.Println("Insert passenger succesfully added... ")
	case 2:
		// Get Passenger By Id
		Passenger, err := BasaPassenger.GetPassengerById()
		if err != nil {
			log.Fatal("error while get passenger by id ", err)
		}
		fmt.Println(Passenger)
	case 3:
		// List of Passengers
		Passengers, err := BasaPassenger.ListOfPassengers()
		if err != nil {
			log.Fatal("error while get lists of passengers ... ", err)
		}
		fmt.Println(Passengers)
	case 4:
		// Update Passenger By Id
		if err := BasaPassenger.UpdatePassengerById(); err != nil {
			log.Fatal("error while update passenger by id ... ", err)
		}
		fmt.Println("passenger updated succesfully...")
	case 5:
		// Delete Passenger By Id
		if err := BasaPassenger.DeletPassengerByID(); err != nil {
			log.Fatal("error while delete passenger")
		}
		fmt.Println("passenger deleted succesfully... ")
	default:
		fmt.Println("no such command exists...")
	}
}
func (b BasaPassengers) InsertPassenger() error {
	psgr := structforpassengers.ForInsertAndUpdatePassenger()
	_, err := b.db.Exec(`insert into passengers (id, first_name, last_name, email, phone, ticket_id) values ($1, $2, $3, $4, $5, $6)`, psgr.ID, psgr.FirstName, psgr.LastName, psgr.Email, psgr.Phone, psgr.TicketID)
	if err != nil {
		return err
	}
	return nil
}
func (b BasaPassengers) GetPassengerById() (structforpassengers.Passengers, error) {
	PID := structforpassengers.ForGetAndDeletePassengerById()
	row := b.db.QueryRow(`select id, first_name, last_name, email, phone, ticket_id from passengers where id = $1`, PID)
	GetPassenger := structforpassengers.Passengers{}
	err := row.Scan(&PID.ID, &PID.FirstName, &PID.LastName, &PID.Email, &PID.Phone, &PID.TicketID)
	if err != nil {
		return structforpassengers.Passengers{}, err
	}
	return GetPassenger, nil
}

func (b BasaPassengers) ListOfPassengers() ([]structforpassengers.Passengers, error) {
	Passengers := []structforpassengers.Passengers{}
	rows, err := b.db.Query(`select id, first_name, last_name, email, phone, ticket_id from passengers`)
	if err != nil {
		return []structforpassengers.Passengers{}, err
	}
	for rows.Next() {
		Passenger := structforpassengers.Passengers{}
		rows.Scan(&Passenger.ID, &Passenger.FirstName, &Passenger.LastName, &Passenger.Email, &Passenger.Phone, &Passenger.TicketID)
		Passengers = append(Passengers, Passenger)
	}
	return Passengers, nil
}

func (b BasaPassengers) UpdatePassengerById() error {
	passenger := structforpassengers.ForInsertAndUpdatePassenger()
	_, err := b.db.Exec(`update passengers set first_name = $1, last_name = $2, email = $3, phone = $4, ticket_id = $5 where id = $6`, passenger.FirstName, passenger.LastName, passenger.Email, passenger.Phone, passenger.TicketID, passenger.ID)
	if err != nil {
		return err
	}
	return nil
}

func (b BasaPassengers) DeletPassengerByID() error {
	deleting := structforpassengers.ForGetAndDeletePassengerById()
	_, err := b.db.Exec(`delete from  passengers where id = $1`, deleting.ID)
	if err != nil {
		return err
	}
	return nil
}
