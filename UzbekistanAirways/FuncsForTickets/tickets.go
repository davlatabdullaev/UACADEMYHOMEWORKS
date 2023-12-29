package funcsfortickets

import (
	connectdb "basa-airways/ConnectDB"
	structfortickets "basa-airways/StructForTickets"
	"database/sql"
	"fmt"
	"log"
)


type BasaTickets struct{
	db *sql.DB
}

func New(db *sql.DB) BasaTickets{
	return BasaTickets{
		db: db,
	}
}

func FuncsForTickets() {
db, err :=	connectdb.Connectdb()
if err!=nil{
	log.Fatal("error while connect db .... ", err)
}
defer db.Close()
BasaTicket := New(db)
command := 0
fmt.Println("select command: ")
fmt.Print(`
  1 - insert ticket
  2 - get ticket by id
  3 - get lists ticket
  4 - update ticket by id
  5 - delete ticket by id
`)
fmt.Scan(&command)
switch command {
	case 1:
	// insert ticket
if err :=	BasaTicket.InsertTicket(); err!=nil{
	log.Fatal("error while insert ticket...", err)
}
fmt.Println("insert ticket completed succesfully...")
case 2: 
// get ticket by id
ticket, err := BasaTicket.GetTicketById()
if err!=nil{
	log.Fatal("error while get ticket by id...", err)
}
fmt.Println(ticket)
case 3: 
// get lists ticket
 GetList, err := BasaTicket.GetTicketLists()
 if err!=nil{
	log.Fatal("error while get ticket lists...", err)
 }
 fmt.Println(GetList)
case 4:
	// update ticket by id
if err :=	BasaTicket.UpdateTicketById(); err!=nil{
	log.Fatal("error while update ticket by id...", err)
}
fmt.Println("update ticket by id completed succesfully...")
case 5: 
// delete ticket by id
err := BasaTicket.DeleteTicketById()
if err!=nil{
	log.Fatal("error while delete ticket by id...", err)
}
fmt.Println("deleted completed succesfully")
default: 
fmt.Println(" no such command exists...")
}
}


func (b BasaTickets) InsertTicket() error{
newticket := structfortickets.InsertAndUpdateTicket()
_, err := b.db.Exec(`insert into tickets (id, from_city, to_city, flight_date) values ($1, $2, $3, $4)`, newticket.ID, newticket.FromCity, newticket.ToCity, newticket.FlightDate)
if err!=nil{
	return err
}
	return nil
}
func (b BasaTickets) GetTicketById() (structfortickets.Tickets, error){
	t := structfortickets.GetAndDeleteTicket()
	row := b.db.QueryRow(`select id, from_city, to_city, flight_date from tickets where id = $1`, t.ID)
	GetTicket := structfortickets.Tickets{}
	err := row.Scan(&GetTicket.ID, &GetTicket.FromCity, &GetTicket.ToCity, &GetTicket.FlightDate)
	if err!=nil{
		return structfortickets.Tickets{}, err
	}
	return GetTicket, nil
}

func (b BasaTickets) GetTicketLists() ([]structfortickets.Tickets, error) {
	GetAllTickets := []structfortickets.Tickets{}
	rows, err := b.db.Query(`select id, from_city, to_city, flight_date from tickets`)
	if err!=nil{
		return nil, err
	}
	for rows.Next() {
		OneTicket := structfortickets.Tickets{}
		rows.Scan(&OneTicket.ID, &OneTicket.FromCity, &OneTicket.ToCity, &OneTicket.FlightDate)
		GetAllTickets= append(GetAllTickets, OneTicket)
	}
	return GetAllTickets, nil
}

func (b BasaTickets) UpdateTicketById() error{
	t := structfortickets.InsertAndUpdateTicket()
	_, err := b.db.Exec(`update tickets set from_city = $1, to_city = $2, flight_date = $3 where id = $4`, t.FromCity, t.ToCity, t.FlightDate, t.ID)
     if err!=nil{
		return err
	 }
	 return nil
}
 func (b BasaTickets) DeleteTicketById() error{
	t := structfortickets.GetAndDeleteTicket()
	_, err := b.db.Exec(`delete from tickets where id = $1`, t.ID)
	if err!=nil{
		return err
	}
	return nil
 }



