package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/config"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/routers"
	"log"
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

	routersInit.Run(":" + conf.ServerConfig.Port)
}
