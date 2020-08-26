package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/config"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/routers"
	"log"
	"os"
)

// init is invoked before main()
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	conf := config.LoadConfiguration()
	routersInit := routers.InitRouter()

	db, err := sql.Open("pgx", conf.DbConfig.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	routersInit.Run(":" + conf.ServerConfig.Port)
}
