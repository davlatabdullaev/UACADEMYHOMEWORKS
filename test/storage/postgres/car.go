package postgres

import (
	"database/sql"
	"test/models"

	"github.com/google/uuid"
)

type carRepo struct {
	DB *sql.DB
}

func NewCarRepo(db *sql.DB) carRepo {
	return carRepo{
		DB: db,
	}
}

func (c carRepo) Insert(car models.Car) (string, error) {
	id := uuid.New().String()
	_, err := c.DB.Exec(`insert into cars values ($1, $2, $3, $4)`, id, car.Model, car.Brand, car.Year)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c carRepo) GetByID(id uuid.UUID) (models.Car, error) {
	car := models.Car{}
	if err := c.DB.QueryRow(`select id, model, brand, year from cars where id = $1`, id).Scan(
		&car.ID,
		&car.Model,
		&car.Brand,
		&car.Year,
	); err != nil {
		return models.Car{}, err
	}
	return car, nil
}

func (c carRepo) GetList() ([]models.Car, error) {
	cars := []models.Car{}
	rows, err := c.DB.Query(`select id, model, brand, year from cars`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		car := models.Car{}
		if err := rows.Scan(&car.ID, &car.Model, &car.Brand, &car.Year); err != nil {
			return nil, err
		}

		cars = append(cars, car)

	}

	return cars, nil
}

func (c carRepo) Update(id uuid.UUID, car models.Car) error {
if _, err	:= c.DB.Exec(`update cars set model = $1, brand = $2, year = $3 where id = $4`, car.Model, car.Brand, car.Year, id); err!=nil{
	return err
}
 return nil
}

func (c carRepo) Delete(id uuid.UUID) error {
	if _, err := c.DB.Exec(`delete from cars where id = $1`, id); err!=nil{
		return err
	}
	return nil
}
