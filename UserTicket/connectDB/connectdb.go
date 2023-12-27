package connectdb

import "database/sql"


func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=davlat password=1 database=airport sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil

	
}


