package repository

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"log"
	"os"
)

var db *sqlx.DB

func Setup() {
	var err error
	db, err = sqlx.Connect("pgx", setting.AppSetting.DbConfig.Url)
	if err != nil {
		log.Fatalln("postgres.Setup err: %v", err)
		os.Exit(1)
	}
}
