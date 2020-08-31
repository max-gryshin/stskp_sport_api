package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/setting"
	"log"
	"os"
)

var db *pgx.Conn

func Setup() {
	var err error
	db, err = pgx.Connect(context.Background(), setting.AppSetting.DbConfig.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		log.Fatalf("models.Setup err: %v", err)
		os.Exit(1)
	}
}

//// CloseDB closes database connection (unnecessary)
//func CloseDB() {
//	defer db.Close()
//}
