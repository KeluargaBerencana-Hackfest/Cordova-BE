package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Ndraaa15/cordova/utils/errors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ClientDB struct {
	*sqlx.DB
}

func ConnDatabase() (*ClientDB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Printf("[cordova-database] failed to connect postgres database. Error : %v\n", err)
		return nil, errors.ErrConnDatabase
	}

	return &ClientDB{db}, nil
}
