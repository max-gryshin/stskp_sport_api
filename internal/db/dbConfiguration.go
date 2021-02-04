package db

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"

	"github.com/jmoiron/sqlx"

	"log"
)

const maxIdleCons = 100
const maxOpenCons = 10

var db sqlx.DB

func GetDB() *sqlx.DB {
	return &db
}

func Setup(settings *setting.Setting) {
	db, err := sqlx.Connect("pgx", settings.DBConfig.URL)
	if err != nil {
		log.Fatalf("postgres.Setup err: %v\n", err)
	}
	db.SetMaxIdleConns(maxIdleCons)
	db.SetMaxOpenConns(maxOpenCons)
}
