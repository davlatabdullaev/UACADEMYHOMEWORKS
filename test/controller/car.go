package controller

import (
	"fmt"
	"log"
	"test/models"
	"time"

	"github.com/google/uuid"
)

func (c Controller) CreateCar() {
	car := getCarInfo()
	if car.Year <= 0 || car.Year < time.Now().Year()+1 {
		fmt.Print("year input is not correct")
		return
	}
	id, err := c.Store.CarStorage.Insert(car)
	if err != nil {
		fmt.Println("error while creating data in inside control", err.Error())
		return
	}
	fmt.Println("id: ", id)
}

func (c Controller) GetCarByID() {
	idStr := ""
	fmt.Println("enter id: ")
	fmt.Scan(&idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}
	car, err := c.Store.CarStorage.GetByID(id)
	if err != nil {
		fmt.Println("error while getting car by id ", err.Error())
		return
	}
	fmt.Println("your car is: ", car)
}

func getCarInfo() models.Car {
	var (
		model, brand string
		year         int
	)

	fmt.Print("enter model and brand: ")
	fmt.Scan(&model, &brand)

	fmt.Print("enter year: ")
	fmt.Scan(&year)

	return models.Car{
		Model: model,
		Brand: brand,
		Year:  year,
	}
}

func (c Controller) GetList() {
	 car, err := c.Store.CarStorage.GetList()
	 if err!=nil{
	fmt.Println("error while get list car ", err.Error())
	return
} 
   fmt.Println(car)
}

func (c Controller) UpdateCarByID() {
	idStr := ""
	fmt.Println("enter id: ")
	fmt.Scan(&idStr)
	car := getCarInfo()

	id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}
	if err := c.Store.CarStorage.Update(id, car); err!=nil {
		log.Fatal("error while update car by id ", err.Error())
	}
	fmt.Println("car updated succesfully ")
}

func (c Controller) DeleteCarByID() {
	idStr := ""
	fmt.Println("enter id: ")
	fmt.Scan(&idStr)

id, err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid err: ", err.Error())
		return
	}
	if err := c.Store.CarStorage.Delete(id); err!=nil{
		log.Fatal(" err while delete car by id ", err.Error())
	} 
fmt.Println("car delete succesfully ")
}