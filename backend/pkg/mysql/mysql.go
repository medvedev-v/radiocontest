package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Vendor   string
	Username string
	Password string
	DBname   string
}

func NewConnection() *Connection {
	var connection Connection

	connection.Vendor = os.Getenv("DB_VENDOR")
	connection.Username = os.Getenv("DB_USER")
	connection.Password = os.Getenv("DB_PASSWORD")
	connection.DBname = os.Getenv("DB_NAME")

	return &connection
}

func Connect() (*sql.DB, error) {
	var connection = NewConnection()

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", connection.Username, connection.Password, connection.DBname)
	db, err := sql.Open(connection.Vendor, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Successfully connected to the database!")

	return db, nil
}
