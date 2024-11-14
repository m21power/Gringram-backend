package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"log"
)

func ConnectDB() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:abate@tcp(localhost:3306)/gringram?parseTime=true")

	if err != nil {

		log.Fatalf("Could not open database: %v", err)

		return nil, err

	}

	if err = db.Ping(); err != nil {

		log.Fatalf("Could not connect to the database: %v", err)

		return nil, err

	}

	return db, nil
}
