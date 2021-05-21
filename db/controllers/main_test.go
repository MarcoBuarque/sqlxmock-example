package controllers

import (
	"os"
	"testing"

	"github.com/MarcoBuarque/sqlxmock-example/db"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	// load env
	godotenv.Load()

	err := db.StartMock(false)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	defer db.MockNode.ExpectClose()

	// db.MockNode.MatchExpectationsInOrder(false)

	retCode := m.Run()

	os.Exit(retCode)
}
