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

	if os.Getenv("UNIT") == "true" || os.Getenv("UNIT") == "" {
		err := db.StartMock(false)
		if err != nil {
			panic(err)
		}
		// db.MockNode.MatchExpectationsInOrder(false)
	} else {
		err := db.Start()
		if err != nil {
			panic(err)
		}
	}
	defer db.Close()

	retCode := m.Run()

	os.Exit(retCode)
}
