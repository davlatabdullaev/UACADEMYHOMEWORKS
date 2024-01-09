package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"test/models"
	"test/pkg/checker"

	"github.com/google/uuid"
)

func (c Controller) Car(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		c.CreateCar(w, r)
	case http.MethodGet:
		c.GetCarByID(w, r)
	case http.MethodPut:
		c.UpdateCarByID(w, r)
	case http.MethodDelete:
		c.DeleteCarByID(w, r)
	}
}

// CREATE

func (c Controller) CreateCar(w http.ResponseWriter, r *http.Request) {
	car := getCarInfo()

	id, err := c.Store.CarStorage.Insert(car)
	if err != nil {
		fmt.Println("error while creating data in inside control", err.Error())
		handResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	handResponse(w, http.StatusCreated, id)
}

// READ BY ID

func (c Controller) GetCarByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		handResponse(w, http.StatusBadRequest, errors.New("missing car ID"))
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		handResponse(w, http.StatusBadRequest, errors.New("invalid car ID format"))
		return
	}

	car, err := c.Store.CarStorage.GetByID(id)
	if err != nil {
		fmt.Println("error while getting car by id ", err.Error())
		handResponse(w, http.StatusInternalServerError, err)
		return
	}

	handResponse(w, http.StatusOK, car)
}

// UPDATE

func (c Controller) UpdateCarByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		handResponse(w, http.StatusBadRequest, errors.New("missing car ID"))
		return
	}

	car := getCarInfo()

	id, err := uuid.Parse(idStr)
	if err != nil {
		handResponse(w, http.StatusBadRequest, errors.New("invalid car ID format"))
		return
	}

	if err := c.Store.CarStorage.Update(id, car); err != nil {
		log.Fatal("error while updating car by id ", err.Error())
		handResponse(w, http.StatusInternalServerError, err)
		return
	}

	handResponse(w, http.StatusOK, "Car updated successfully")
}

// DELETE

func (c Controller) DeleteCarByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		handResponse(w, http.StatusBadRequest, errors.New("missing car ID"))
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		handResponse(w, http.StatusBadRequest, errors.New("invalid car ID format"))
		return
	}

	if err := c.Store.CarStorage.Delete(id); err != nil {
		log.Fatal("error while deleting car by id ", err.Error())
		handResponse(w, http.StatusInternalServerError, err)
		return
	}

	handResponse(w, http.StatusOK, "Car deleted successfully")
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
