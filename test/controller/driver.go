package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"test/models"
	"test/pkg/checker"
	"time"

	"github.com/google/uuid"
)

func (c Controller) Driver(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateDriver(w, r)
	case http.MethodGet:
		values := r.URL.Query()
		fmt.Println("values: ", values, time.Now())
		_, ok := values["id"]
		if ok {
			c.GetDriverByID(w, r)
		} else {
			c.GetDriversList(w, r)
		}
	case http.MethodPut:
		c.UpdateDriverByID(w, r)
	case http.MethodDelete:
		c.DeleteDriverByID(w, r)
	}
}

// CREATE

func (c Controller) CreateDriver(w http.ResponseWriter, r *http.Request) {
	driver := models.Driver{}

	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		log.Fatal("error while reading data from client", err.Error())
		handResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !checker.PhoneNumber(driver.Phone) {
		fmt.Println("the phone number format is not correct!")
		handResponse(w, http.StatusBadRequest, errors.New("phone type is not correct!"))
		return
	}

	id, err := c.Store.DriverStorage.InsertDriver(driver)
	if err != nil {
		fmt.Println("error while inserting driver inside controller err: ", err.Error())
		handResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	handResponse(w, http.StatusCreated, id)
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

// READ BY ID

func (c Controller) GetDriverByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	fmt.Println("values : ", values)
	id := values["id"][0]
	driver, err := c.Store.DriverStorage.GetByID(id)
	if err != nil {
		fmt.Println("error while getting driver by id ", err.Error())
		handResponse(w, http.StatusInternalServerError, err)
		return
	}

	handResponse(w, http.StatusOK, driver)

}

// READ ALL

func (c Controller) GetDriversList(w http.ResponseWriter, r *http.Request) {
	drivers, err := c.Store.DriverStorage.GetListDriver()
	if err != nil {
		fmt.Println("error while get list driver ", err.Error())
		handResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(drivers)
}

// UPDATE

func (c Controller) UpdateDriverByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values.Get("id")
	if id == "" {
		handResponse(w, http.StatusBadRequest, errors.New("missing driver ID"))
		return
	}

	driver := models.Driver{}
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		log.Fatal("error while reading data from client", err.Error())
		handResponse(w, http.StatusBadRequest, err.Error())
		return
	}


     uid:= uuid.MustParse(id)
	err := c.Store.DriverStorage.UpdateDriverByID(uid, driver)
	if err != nil {
		fmt.Println("error while updating driver by ID inside controller err: ", err.Error())
		handResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	handResponse(w, http.StatusOK, "Driver updated successfully")

}

// DELETE

func (c Controller) DeleteDriverByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id := values.Get("id")
	if id == "" {
		handResponse(w, http.StatusBadRequest, errors.New("missing driver ID"))
		return
	}

	uid := uuid.MustParse(id)
	err := c.Store.DriverStorage.DeleteDriverByID(uid)
	if err != nil {
		fmt.Println("error while deleting driver by ID inside controller err: ", err.Error())
		handResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	handResponse(w, http.StatusOK, "Driver deleted successfully")

}
