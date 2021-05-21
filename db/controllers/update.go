package controllers

import (
	"fmt"

	"github.com/MarcoBuarque/sqlxmock-example/db"
)

func UpdateUserLastName(userID int32, lastName string) (bool, error) {

	tx := db.Conn.MustBegin().Tx
	defer tx.Rollback()

	res, err := tx.Exec(`UPDATE users SET last_name=$1 WHERE id=$2`, lastName, userID)
	if err != nil {
		tx.Rollback()
		return false, fmt.Errorf("%v: %v", StrInsertUserFail, err)
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return false, fmt.Errorf("%v: %v", StrInsertUserFail, commitErr)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("%v: %v", StrInsertUserFail, err)
	}

	return rowsAffected > 0, nil
}
