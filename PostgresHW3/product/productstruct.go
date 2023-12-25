package product

import "time"

type Product struct {
	ID         int
	Name       string
	Price      float64
	CategoryID int
	CreateAt   time.Time
	UpdateAt   time.Time
}
