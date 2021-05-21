package controllers

import (
	"fmt"

	"github.com/MarcoBuarque/sqlxmock-example/db"
	"github.com/MarcoBuarque/sqlxmock-example/db/models"
)

func SelectUsers() ([]models.User, error) {
	var users []models.User

	err := db.Conn.Select(&users, "SELECT id, first_name, last_name, email, country, created_at FROM users")

	if err != nil {
		return users, fmt.Errorf("%v: %v", StrGetUsersFail, err)
	}

	if len(users) == 0 {
		return users, fmt.Errorf("%v: %v", StrGetUsersFail, ErrEmptyUsers)
	}

	return users, nil
}

func SelectUserByID(userID int32) (models.User, error) {
	var user models.User

	err := db.Conn.Get(&user, "SELECT id, first_name, last_name, email, country, created_at FROM users WHERE id = $1", userID)

	if err != nil {
		return user, fmt.Errorf("%v: %v", StrGetUserFail, err)
	}

	return user, nil
}
