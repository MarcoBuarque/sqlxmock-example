package controllers

import (
	"fmt"
	"testing"

	"github.com/MarcoBuarque/sqlxmock-example/db"
	"github.com/MarcoBuarque/sqlxmock-example/db/models"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestInsertUser(t *testing.T) {
	assert := assert.New(t)

	// If we change the match option to QueryMatcherEqual, we need to use the second option
	query := "INSERT INTO users"
	// query := "INSERT INTO users (first_name, last_name, email, country) VALUES ($1, $2, $3, $4) RETURNING id"

	columns := []string{"id"}

	user := models.User{
		FristName: "Marco",
		LastName:  "Oliveira",
		Email:     "test@test.com",
		Country:   "Brazil",
	}
	tests := []struct {
		purpose   string
		giveUser  models.User
		mock      func()
		wantID    int32
		wantError error
	}{
		{
			purpose:   "Test Invalid args",
			giveUser:  models.User{},
			mock:      func() {},
			wantID:    0,
			wantError: fmt.Errorf("%v, %v", StrInsertUserFail, ErrParameterNotFound),
		},
		{
			purpose:  "Test Success",
			giveUser: user,
			mock: func() {
				rows := sqlxmock.NewRows(columns).AddRow(1)

				db.MockNode.ExpectBegin() // TODO: teste do anyargument (insert user created at)
				db.MockNode.ExpectQuery(query).WithArgs(user.FristName, user.LastName, user.Email, user.Country).WillReturnRows(rows)
				db.MockNode.ExpectCommit()
			},
			wantID:    1,
			wantError: nil,
		},
		{
			purpose:  "Should return an error from db",
			giveUser: user,
			mock: func() {
				db.MockNode.ExpectBegin()
				db.MockNode.ExpectQuery(query).WithArgs(user.FristName, user.LastName, user.Email, user.Country).WillReturnError(fmt.Errorf("um erro qualquer"))
				// db.MockNode.ExpectCommit()
				db.MockNode.ExpectRollback()

			},
			wantID:    0,
			wantError: fmt.Errorf("%v: %v", StrInsertUserFail, "um erro qualquer"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.purpose, func(t *testing.T) {
			tt.mock()
			id, err := InsertUser(tt.giveUser)

			assert.Equal(tt.wantError, err)
			assert.Equal(tt.wantID, id)
		})
	}

	assert.Equal(nil, db.MockNode.ExpectationsWereMet())
}
