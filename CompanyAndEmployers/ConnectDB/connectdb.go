package connectdb

import (
	_"github.com/lib/pq"
	"database/sql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=davlat password=1 database=companies sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
