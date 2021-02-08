package db

import (
	_ "github.com/jackc/pgx/stdlib" // need to connect with db
	"github.com/jmoiron/sqlx"
	"log"
)

const maxIdleCons = 100
const maxOpenCons = 10

var DB sqlx.DB

func CreateDBConnection(url string) *sqlx.DB {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatalf("postgres.Setup err: %v\n", err)
	}
	DB = *db
	DB.SetMaxIdleConns(maxIdleCons)
	DB.SetMaxOpenConns(maxOpenCons)

	return &DB
}
