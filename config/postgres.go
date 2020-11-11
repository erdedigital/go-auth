package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func PostgresDB() (*sql.DB, error) {

	pgURL := fmt.Sprintf("%s", os.Getenv("POSTGRES_URL"))
	db, err := CreateConnection(pgURL)
	if err != nil {
		panic(err)
	}
	return db, nil
}

func CreateConnection(pgURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", pgURL)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(100)
	db.SetMaxOpenConns(100)

	return db, nil
}
