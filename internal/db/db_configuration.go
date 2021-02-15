package db

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	_ "github.com/jackc/pgx/stdlib" // need to connect with db

	"github.com/jmoiron/sqlx"

	"log"
	"strconv"
)

var DB sqlx.DB

func CreateDBConnection(dbSetting *setting.DBSetting) *sqlx.DB {
	db, err := sqlx.Connect("pgx", dbSetting.URL)
	if err != nil {
		log.Fatalf("postgres.Setup err: %v\n", err)
	}
	DB = *db
	maxIdleCons, err := strconv.Atoi(dbSetting.MaxIdleCons)
	if err != nil {
		log.Fatalf("postgres.Setup err: %v\n", err)
	}
	maxOpenCons, err := strconv.Atoi(dbSetting.MaxOpenCons)
	if err != nil {
		log.Fatalf("postgres.Setup err: %v\n", err)
	}
	DB.SetMaxIdleConns(maxIdleCons)
	DB.SetMaxOpenConns(maxOpenCons)

	return &DB
}
