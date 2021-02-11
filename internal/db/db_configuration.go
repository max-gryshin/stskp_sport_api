package db

import (
	_ "github.com/jackc/pgx/stdlib" // need to connect with db

	"github.com/jmoiron/sqlx"

	"log"
)

const MaxIdleCons = 100
const MaxOpenCons = 10

var DB sqlx.DB

func CreateDBConnection(url string) *sqlx.DB {
	db, err := sqlx.Connect("pgx", url)
	if err != nil {
		log.Fatalf("postgres.Setup err: %v\n", err)
	}
	DB = *db
	DB.SetMaxIdleConns(MaxIdleCons)
	DB.SetMaxOpenConns(MaxOpenCons)

	return &DB
}
