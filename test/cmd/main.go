package main

import (
	"fmt"
	"log"
	"net/http"
	"test/config"
	"test/controller"
	"test/storage/postgres"
)

func main() {

	cfg := config.Load()

	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db err: ", err)
		return
	}
	defer store.DB.Close()

	con := controller.New(store)

	//http.HandleFunc("/driver", con.Driver)
	http.HandleFunc("/car", con.Car)
    fmt.Println("listening at: ")
	http.ListenAndServe(":8080", nil)
}
