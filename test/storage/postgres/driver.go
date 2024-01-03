package postgres

import (
	"database/sql"
	"test/models"
	"github.com/google/uuid"
)

type driverRepo struct{
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) driverRepo{
	return driverRepo{
		DB: db,
	}
}


func (d driverRepo) InsertDriver(driver models.Driver) (string, error) {
   id := uuid.New()

	if _, err := d.DB.Exec(`insert into drivers values ($1, $2, $3)`, id, driver.FullName, driver.Phone); err!= nil{
		return "", err
	}
	return id.String(), nil
}

func (d driverRepo) GetDriverByID(id uuid.UUID) (models.Driver, error) {
	driver := models.Driver{}
	if err := d.DB.QueryRow(`select id, model, brand, year from cars where id = $1`, id).Scan(
		&driver.ID,
		&driver.FullName,
		&driver.Phone,
	); err != nil {
		return models.Driver{}, err
	}
	return driver, nil
}

func (d driverRepo) GetListDriver() ([]models.Driver, error) {
	drivers := []models.Driver{}
	rows, err := d.DB.Query(`select id, model, brand, year from cars`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		driver := models.Driver{}
		if err := rows.Scan(&driver.ID, &driver.FullName, &driver.Phone); err != nil {
			return nil, err
		}

		drivers = append(drivers, driver)

	}

	return drivers, nil
}

func (d driverRepo) UpdateDriverByID(id uuid.UUID, driver models.Driver) error {
if _, err	:= d.DB.Exec(`update drivers set full_name = $1, phone = $2 where id = $3`, driver.FullName, driver.Phone, id); err!=nil{
	return err
}
 return nil
}

func (d driverRepo) DeleteDriverByID(id uuid.UUID) error {
	if _, err := d.DB.Exec(`delete drivers where id = $1`, id); err!=nil{
		return err
	}
	return nil
}
