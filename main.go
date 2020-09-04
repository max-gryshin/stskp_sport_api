package main

import (
	"github.com/joho/godotenv"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/routers"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"log"
)

var conf *setting.Setting

// init is invoked before main()
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	conf = setting.LoadSetting()
	logging.Setup()
	repository.Setup()
}

func main() {
	routersInit := routers.InitRouter()
	routersInit.Run(":" + conf.ServerConfig.Port)
}
