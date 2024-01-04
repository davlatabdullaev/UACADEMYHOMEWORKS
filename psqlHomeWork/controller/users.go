package controller

import (
	"basa/structs"
	"fmt"
	"log"
	"regexp"

	"github.com/google/uuid"
)

func (c Controller) InsertUser() {
	user := getUserInfo()

	phoneRegex := regexp.MustCompile(`^\+\d{12}$`)
	if !phoneRegex.MatchString(user.Phone) {
		fmt.Println("phone number is not correct ")
		return
	}

	id, err := c.Store.UsersStorage.InsertUser(user)
	if err != nil {
		fmt.Println("error while creating data inside controller err: ", err.Error())
		return
	}
	fmt.Println(id)
}
func (c Controller) GetUserByID() {
	idStr := ""
	fmt.Print("enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}
	c.Store.UsersStorage.GetByIDUser(id)
}

func getUserInfo() structs.Users {
	var (
		first_name string
		last_name  string
		email      string
		phone      string
	)
	fmt.Println("enter first name: ")
	fmt.Scan(&first_name)
	fmt.Println("enter last name: ")
	fmt.Scan(&last_name)
	fmt.Println("enter email: ")
	fmt.Scan(&email)
	fmt.Println("enter phone: ")
	fmt.Scan(&phone)
	return structs.Users{
		First_Name: first_name,
		Last_Name:  last_name,
		Email:      email,
		Phone:      phone,
	}
}

func (c Controller) UpdateUser() {
user := getUserInfoForUpdate()
c.Store.UsersStorage.UpdateUser(user)

}
func (c Controller) DeleteUser() {
var	idStr string 
fmt.Println("enter id: ")
fmt.Scan(&idStr)
id, err := uuid.Parse(idStr)
if err!=nil{
	log.Fatalln("this id is not uuid type err: ", err.Error())
}
c.Store.UsersStorage.DeleteUser(id)
}

func (c Controller) GetUsersList() {
Users, err :=	c.Store.UsersStorage.GetListUser()
if err!=nil{
	log.Fatal("error while get users list ", err.Error())
}
fmt.Println(Users)
}



func  getUserInfoForUpdate() structs.Users{
var idStr string
fmt.Println("Enter id: ")
fmt.Scan(&idStr)
var (
	first_name string
	last_name  string
	email      string
	phone      string
)
id, err := uuid.Parse(idStr)
if err!=nil{
	fmt.Println("id is not uuid err: ", err.Error())
}
fmt.Println("enter first name: ")
fmt.Scan(&first_name)
fmt.Println("enter last name: ")
fmt.Scan(&last_name)
fmt.Println("enter email: ")
fmt.Scan(&email)
fmt.Println("enter phone: ")
fmt.Scan(&phone)
return structs.Users{
	ID: id,
	First_Name: first_name,
	Last_Name: last_name,
	Email: email,
	Phone: phone,
}

}