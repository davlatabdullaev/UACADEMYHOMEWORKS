package inventory

import (
	"database/sql"
	"go-backend/country"
)
type Inventory struct{
	db *sql.DB
}

func New(db *sql.DB) Inventory{
	return Inventory{
		db: db,
	}
}

func (i Inventory) InsertCountry(c country.Country) error{
	if _, err := i.db.Exec(
		`insert into countries values ($1, $2, $3)`, c.ID, c.Name, c.Currency,
	); err!=nil{
     return err
	}
	return nil
}
func (i Inventory) GetAllCountries() ([]country.Country, error) {
	rows, err := i.db.Query(
		`select * from countries`,
	)
	if err!=nil{
		return nil, err
	}
	cs := []country.Country{}

	for rows.Next(){
		c:=country.Country{}
		if err = rows.Scan(&c.ID, &c.Name, &c.Currency); err!=nil{
			return nil, err
		}
		cs=append(cs, c)
	}
	return cs, nil
}
func (i Inventory) GetCountryByID(id string) (country.Country, error) {
	 row := i.db.QueryRow(`select id, name, currency from countries where id = $1`, id)
	
	 m := country.Country{}

	if err := row.Scan(&m.ID, &m.Name, &m.Currency); err!=nil{
		panic(err)
	}
	
	return country.Country{}, nil
}

func (i Inventory) DeleteCountry(id string) error {
   if _, err := i.db.Exec(`delete from countries where id = $1`, id); err!=nil{
	panic(err)
   }
	return nil
}
