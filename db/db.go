package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "aldizix6ZWY"
	dbname   = "assignmentsDua"
)

var (
	db  *sql.DB
	err error
)

func InitializeDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	createOrdersTable := `
		CREATE TABLE IF NOT EXISTS orders (
			order_id SERIAL PRIMARY KEY,
			customer_name VARCHAR(255) NOT NULL,
			ordered_at TIMESTAMPTZ DEFAULT now()
		)
	`
	createItemsTable := `
		CREATE TABLE IF NOT EXISTS items (
			item_id SERIAL PRIMARY KEY,
			order_id INT NOT NULL,
			item_code VARCHAR(255) NOT NULL,
			description VARCHAR(255) NOT NULL,
			quantity INT NOT NULL
		)
	`

	createTable(createOrdersTable)
	createTable(createItemsTable)

	fmt.Println("Successfully connected to the database!")
}

func createTable(tableName string) {
	_, err = db.Exec(tableName)
	if err != nil {
		log.Fatal("Error creating table:", err.Error())
	}
}

func GetDB() *sql.DB {
	return db
}
