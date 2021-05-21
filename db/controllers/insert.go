package controllers

import (
	"fmt"

	"github.com/MarcoBuarque/sqlxmock-example/db"
	"github.com/MarcoBuarque/sqlxmock-example/db/models"
)

func InsertUser(user models.User) (int32, error) {
	if user.FristName == "" || user.Email == "" || user.Country == "" {
		return 0, fmt.Errorf("%v, %v", StrInsertUserFail, ErrParameterNotFound)
	}

	var id int32

	tx := db.Conn.MustBegin().Tx
	defer tx.Rollback()

	err := tx.QueryRow(`INSERT INTO users (first_name, last_name, email, country) VALUES ($1, $2, $3, $4) RETURNING id`,
		user.FristName, user.LastName, user.Email, user.Country).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("%v: %v", StrInsertUserFail, err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return 0, fmt.Errorf("%v: %v", StrInsertUserFail, commitErr)
	}

	return id, nil
}
