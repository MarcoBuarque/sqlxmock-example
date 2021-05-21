package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

var Conn *sqlx.DB

var MockNode sqlxmock.Sqlmock

func Start() error {
	fmt.Println("Connecting to DB")

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "postgres"
	}

	db, err := sqlx.Connect(dbDriver, GetURI())

	if err != nil {
		fmt.Printf("Something is wrong: %v", err)
		log.Fatalln(err)
		return err
	}

	fmt.Println("Successfully connected!")
	Conn = db
	return nil
}

func GetURI() string {
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_URI")
	port := os.Getenv("DB_PORT")
	userName := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, userName, password, dbName)

	return psqlInfo
}

func Close() {
	Conn.Close()
}

func Ping() error {
	return Conn.Ping()
}

func StartMock(matcherEqual bool) error {
	fmt.Println("Start mock DB")

	var db *sqlx.DB
	var mock sqlxmock.Sqlmock
	var err error

	if matcherEqual {
		db, mock, err = sqlxmock.Newx(sqlxmock.QueryMatcherOption(sqlxmock.QueryMatcherEqual))
	} else {
		db, mock, err = sqlxmock.Newx()
	}
	if err != nil {
		fmt.Printf("Something is wrong: %v", err)
		log.Fatalln(err)
		return err
	}

	fmt.Println("Successfully mocked!")
	Conn = db
	MockNode = mock

	return nil
}
