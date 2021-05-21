package controllers

import (
	"testing"

	"github.com/MarcoBuarque/sqlxmock-example/db"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestUpdateUserLastName(t *testing.T) {
	assert := assert.New(t)

	// If we change the match option to QueryMatcherEqual, we need to use the second option
	query := "UPDATE users"
	// query := "UPDATE users SET last_name=$1 WHERE id=$2"

	tests := []struct {
		purpose      string
		giveUserID   int32
		giveLastName string
		mock         func()
		wantResponse bool
		wantError    error
	}{
		{
			purpose:      "Test Success",
			giveUserID:   57,
			giveLastName: "TEST_LAST_NAME",
			mock: func() {
				db.MockNode.ExpectBegin()
				db.MockNode.ExpectExec(query).WithArgs("TEST_LAST_NAME", 57).WillReturnResult(sqlxmock.NewResult(57, 1))
				db.MockNode.ExpectCommit()
			},
			wantResponse: true,
			wantError:    nil,
		},
		{
			purpose:      "Test invalid userID",
			giveUserID:   99,
			giveLastName: "TEST_INVALID_ID",
			mock: func() {
				db.MockNode.ExpectBegin()
				db.MockNode.ExpectExec(query).WithArgs("TEST_INVALID_ID", 99).WillReturnResult(sqlxmock.NewResult(0, 0))
				db.MockNode.ExpectCommit()
			},
			wantResponse: false,
			wantError:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.purpose, func(t *testing.T) {
			tt.mock()
			updated, err := UpdateUserLastName(tt.giveUserID, tt.giveLastName)

			assert.Equal(tt.wantError, err)
			assert.Equal(tt.wantResponse, updated)
		})
	}

	assert.Equal(nil, db.MockNode.ExpectationsWereMet())
}
