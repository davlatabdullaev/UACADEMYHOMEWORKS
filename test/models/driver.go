package models

import "github.com/google/uuid"

type Driver struct {
	ID       uuid.UUID   `json:"_id"`
	FullName string      `json:"full_name"`
	Phone    string      `json:"phone"`
}





