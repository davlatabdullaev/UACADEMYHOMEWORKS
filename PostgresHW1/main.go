package main

import (
	"database/sql"
	//"fmt"
     "time"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time	
}

func main() {
	db, err := sql.Open("postgres",
		"host = localhost port=5432 user = davlat password=1 database = products sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	products := []Product{
		{ID: uuid.New(), Name: "Burger", Description: "Burger", Price: 10.99, },
		{ID: uuid.New(), Name: "Pizza", Description: "Tasty", Price: 15.99},
		{ID: uuid.New(), Name: "Salad", Description: "Healthy", Price:5008.99},
		{ID: uuid.New(), Name: "Sushi", Description: "Exotic", Price: 20.99},
		{ID: uuid.New(), Name: "Pasta", Description: "Italian", Price: 12.99},
		{ID: uuid.New(), Name: "Steak", Description: "Grilled", Price: 6025.99},
		{ID: uuid.New(), Name: "Sandwich", Description: "Gourmet", Price: 18.99},
		{ID: uuid.New(), Name: "Soup", Description: "Spicy", Price: 14.99},
		{ID: uuid.New(), Name: "Smoothie", Description: "Refreshing", Price: 9.99},
		{ID: uuid.New(), Name: "Taco", Description: "Mexican", Price: 6002.99},
		{ID: uuid.New(), Name: "Dumplings", Description: "Dumplings with Love", Price: 16.99},
		{ID: uuid.New(), Name: "Fried Chicken", Description: "Crispy Fried", Price: 11.99},
		{ID: uuid.New(), Name: "Ice Cream", Description: "Sweet", Price: 27.99},
		{ID: uuid.New(), Name: "Coffee", Description: "Aromatic", Price: 13.99},
		{ID: uuid.New(), Name: "Cupcake", Description: "Decadent", Price: 5019.99},
	}
	for _, product := range products {
		
		_, err := db.Exec(`
			INSERT INTO products (id, name, description, price, created_at, updated_at)
			VALUES ($1, $2, $3, $4, current_timestamp, current_timestamp)
		`, product.ID, product.Name, product.Description, product.Price,)

		if err != nil {
			panic(err)
		}
	}
}
