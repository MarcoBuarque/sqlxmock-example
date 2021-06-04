package models

import (
	"time"
)

// User struct
type User struct {
	ID        int32     `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Country   string    `db:"country"`
	CreatedAt time.Time `db:"created_at"`
}
