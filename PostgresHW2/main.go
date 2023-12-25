package main

import (
	"database/sql"
	_"fmt"
	_"github.com/lib/pq"
	"go-backend/inventory"
	_"github.com/google/uuid"
	_"go-backend/country"
)


func main(){
	db, err := connectDB()
	if err != nil{
		panic(err)
	}
	defer db.Close()

	inv:=inventory.New(db)


CountryId := "ccce1ca6-1cbc-4ddf-ada9-2863ef871ac3"


 if err = inv.DeleteCountry(CountryId); err!=nil{
	panic(err)
 }



	// NewCountry := country.Country{
	// 	ID: uuid.New(),
	// 	Name: "Turkiye",
	// 	Currency: "lira",
	// }

	// if err := inv.InsertCountry(NewCountry); err!=nil{
	// 	panic(err)
	// }
	// fmt.Println(NewCountry)

// 	 var parol string= "398053aa-c1cc-4c24-bdea-ae3fb2f7efc5"

//    	countries, err := inv.GetCountryByID(parol)
// 	if err!=nil{
// 		fmt.Println(err)
// 		return
// 	}
//   fmt.Println(countries)

	// countries, err :=inv.GetAllCountries()
	// if err!=nil{
	// 	panic(err)
	// }
	// fmt.Println("all the countries: ", countries)

}
func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=abdullaev password=1 database=mamlakatlar sslmode=disable")
if err!=nil{
	return nil, err
}
return db, nil
}