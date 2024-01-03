package controller

import (
	"fmt"
	"log"
	"test/models"

	"github.com/google/uuid"
)

func (c Controller) CreateDriver() {
	driver := getDriverInfo()

	id, err := c.Store.DriverStorage.InsertDriver(driver)
	if err != nil {
		fmt.Println("error while inserting driver inside controller err: ", err.Error())
		return
	}
	fmt.Println("your new driver's id is: ", id)

}

func getDriverInfo() models.Driver {
	var (
		fullName, phone string
	)

	fmt.Print("enter driver's fullname : ")
	fmt.Scan(&fullName)

	fmt.Print("enter driver's phone : ")
	fmt.Scan(&phone)

	return models.Driver{
		FullName: fullName,
		Phone:    phone,
	}
}

func (c Controller) GetDriverByID() {
	idStr := ""
	fmt.Println("enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}
	driver, err := c.Store.DriverStorage.GetDriverByID(id)
	if err != nil {
		fmt.Println("error while getting driver by id ", err.Error())
		return
	}
	fmt.Println("driver is: ", driver)
}

func (c Controller) GetDriversList() {
	driver, err := c.Store.DriverStorage.GetListDriver()
	if err != nil {
		fmt.Println("error while get list driver ", err.Error())
		return
	}
	fmt.Println(driver)
}

func (c Controller) UpdateDriverByID() {
	idStr := ""
	fmt.Println("enter id: ")
	fmt.Scan(&idStr)
	driver := getDriverInfo()

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}
	if err := c.Store.DriverStorage.UpdateDriverByID(id, driver); err != nil {
		log.Fatal("error while update driver by id ", err.Error())
	}
	fmt.Println("driver updated succesfully ")
}

func (c Controller) DeleteDriverByID() {
	idStr := ""
	fmt.Println("enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}
	if err := c.Store.DriverStorage.DeleteDriverByID(id); err != nil {
		log.Fatal(" err while delete driver by id ", err.Error())
	}
	fmt.Println("driver delete succesfully ")
}
