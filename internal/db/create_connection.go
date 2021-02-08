package db

import (
	"github.com/jmoiron/sqlx"

	"log"
)

type ConnectionSettions struct {
	URL         string
	MaxIdleCons int
	MaxOpenCons int
}

func CreateConnection(setting ConnectionSettions) *sqlx.DB {
	db, err := sqlx.Connect("pgx", setting.URL)
	if err != nil {
		log.Fatalf("postgres.Setup err: %v\n", err)
	}

	db.SetMaxIdleConns(setting.MaxIdleCons)
	db.SetMaxOpenConns(setting.MaxIdleCons)

	return db
}
