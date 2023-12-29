package connectdb

import (
	"database/sql"
	_"github.com/lib/pq"
)

func Connectdb() (*sql.DB, error) {
	var connect string = `
  host=localhost
  port=5432
  user=davlat 
  password=1
  database=airways
  sslmode=disable
  `
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}
	return db, nil
}
