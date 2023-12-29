package structforpassengers

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type Passengers struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Phone     string
	TicketID  uuid.UUID
}

func ForInsertAndUpdatePassenger() Passengers {
	var StrId string
	fmt.Println("Enter passenger id: ")
	fmt.Scan(&StrId)
    var FName string
	fmt.Println("Enter passenger first name : ")
	fmt.Scan(&FName)
	var LName string
	fmt.Println("Enter passenger last name : ")
	fmt.Scan(&LName)
	var PEmail string
	fmt.Println("Enter passenger email : ")
	fmt.Scan(&PEmail)
	var PPhone string
	fmt.Println("Enter passenger phone : ")
	fmt.Scan(&PPhone)
	var PTicketId string
	fmt.Println("Enter passenger ticket id : ")
	fmt.Scan(&PTicketId)
	PID , err := uuid.Parse(StrId)
	if err!=nil{
		log.Fatal("error while uuid parse passengers id ... ", err)
	}
    PTID , err := uuid.Parse(PTicketId)
	if err!=nil{
		log.Fatal("error while uuid parse passenger's ticket id... ", err)
	}
	return Passengers{
		ID: PID,
		FirstName: FName,
		LastName: LName,
		Email: PEmail,
		Phone: PPhone,
		TicketID: PTID,
	}
}

func ForGetAndDeletePassengerById() Passengers{
var	PID string
fmt.Println("enter passenger id: ")
fmt.Scan(&PID) 
PUUID := uuid.MustParse(PID)
return Passengers{
	ID: PUUID,
}
}

