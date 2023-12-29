package structfortickets

import (
	"fmt"

	"github.com/google/uuid"
)

type Tickets struct {
	ID         uuid.UUID
	FromCity   string
	ToCity     string
	FlightDate string
}

func InsertAndUpdateTicket() Tickets {
	var Id string
	fmt.Println("enter ticket id: ")
	fmt.Scan(&Id)
	var fromcity string
	fmt.Println("enter from city: ")
	fmt.Scan(&fromcity)
	var tocity string
	fmt.Println("enter to city: ")
	fmt.Scan(&tocity)
	var flightdate string
	fmt.Println("enter flight date: ")
	fmt.Scan(&flightdate)
	Tuuid := uuid.MustParse(Id)
	return Tickets{
		ID: Tuuid,
		FromCity: fromcity,
		ToCity: tocity,
		FlightDate: flightdate,
	}
}
func GetAndDeleteTicket() Tickets{
	var Id string
	fmt.Println("enter ticket id: ")
	fmt.Scan(&Id)
	Tuuid := uuid.MustParse(Id)
	return Tickets{
		ID: Tuuid,
	}
}