package domain

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
