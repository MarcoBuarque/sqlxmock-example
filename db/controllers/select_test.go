package controllers

import (
	"fmt"
	"testing"
	"time"

	"github.com/MarcoBuarque/sqlxmock-example/db"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestSelectUsers(t *testing.T) {
	assert := assert.New(t)

	// If we change the match option to QueryMatcherEqual, we need to use the second option
	query := "SELECT (.+) FROM users"
	// query := "SELECT id, first_name, last_name, email, country, created_at FROM users"

	columms := []string{"id", "first_name", "last_name", "email", "country", "created_at"}

	tests := []struct {
		purpose   string
		mock      func()
		wantLen   int
		wantError error
	}{
		{
			purpose: "Test Success",
			mock: func() {
				rows := sqlxmock.NewRows(columms).
					AddRow(1, "Marco", "", "test_email", "Brazil", time.Now()).
					AddRow(2, "User", "Test", "test_2_email", "Brazil", time.Now())

				db.MockNode.ExpectQuery(query).WillReturnRows(rows)
			},
			wantLen:   2,
			wantError: nil,
		},
		{
			purpose: "Should return an error when the connection fail",
			mock: func() {
				db.MockNode.ExpectQuery(query).WillReturnError(fmt.Errorf("Erro no DB"))
			},
			wantLen:   0,
			wantError: fmt.Errorf("%v: %v", StrGetUsersFail, "Erro no DB"),
		},
		// { // TODO: erro se não tiver ninguem no db
		// 	purpose: "Should return an error when users table is empty",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.purpose, func(t *testing.T) {
			tt.mock()
			users, err := SelectUsers()

			assert.Equal(tt.wantError, err)
			assert.Equal(tt.wantLen, len(users))
		})
	}

	assert.Equal(nil, db.MockNode.ExpectationsWereMet())
}

func TestSelectUserByIDs(t *testing.T) {
	assert := assert.New(t)

	// If we change the match option to QueryMatcherEqual, we need to use the second option
	query := "SELECT (.+) FROM users"
	// query := "SELECT id, first_name, last_name, email, country, created_at FROM users WHERE id = $1"

	columms := []string{"id", "first_name", "last_name", "email", "country", "created_at"}

	tests := []struct {
		purpose   string
		giveID    int32
		mock      func()
		wantError error
	}{
		{
			purpose: "Test Success",
			giveID:  1,
			mock: func() {
				rows := sqlxmock.NewRows(columms).
					AddRow(1, "Marco", "", "test_email", "Brazil", time.Now())

				db.MockNode.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
			},
			wantError: nil,
		},
		{
			purpose: "Test Fail",
			giveID:  0,
			mock: func() {
				rows := sqlxmock.NewRows(columms)

				db.MockNode.ExpectQuery(query).WithArgs(0).WillReturnRows(rows)
			},
			wantError: fmt.Errorf("%v: %v", StrGetUserFail, "sql: no rows in result set"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.purpose, func(t *testing.T) {
			tt.mock()
			user, err := SelectUserByID(tt.giveID)

			assert.Equal(tt.wantError, err)
			assert.Equal(tt.giveID, user.ID)
		})
	}

	assert.Equal(nil, db.MockNode.ExpectationsWereMet())
}
