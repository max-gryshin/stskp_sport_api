package routers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/db"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

// RegisterAPIV1 initialize routing information
func RegisterAPIV1(router gin.IRoutes, conf *setting.Setting) {
	connection := db.CreateConnection(db.ConnectionSettions{
		URL:         conf.DBConfig.URL,
		MaxIdleCons: 100,
		MaxOpenCons: 10,
	})

	queryBuilder := goqu.Dialect("postgres")
	userRepo := repository.NewUserRepository(connection, queryBuilder)
	userController := controllers.NewUserController(userRepo)

	router.GET("/users", userController.GetUsers)
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUserByID)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
