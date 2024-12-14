package db

import (
	"log"

	"github.com/Tanakaryuki/go-restapi/pkg/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New() *sqlx.DB {
	db, err := Init()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Init() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.DB_DSN)
	if err != nil {
		return nil, err
	}
	return db, nil
}
