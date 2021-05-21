package main

import (
	"github.com/MarcoBuarque/sqlxmock-example/db"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	godotenv.Load()

	err := db.Start()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
