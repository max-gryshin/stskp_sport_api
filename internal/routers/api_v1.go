package routers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/db"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	v1 "github.com/ZmaximillianZ/stskp_sport_api/internal/routers/api/v1"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

// RegisterAPIV1 initialize routing information
func RegisterAPIV1(router gin.IRoutes, conf *setting.Setting) {
	queryBuilder := goqu.Dialect("postgres")
	userRepo := repository.NewUserRepository(&db.DB, queryBuilder)
	userController := controllers.NewUserController(userRepo)

	router.POST("/api/user/auth", userController.GetAuth)
	router.POST("/api/user/create", v1.CreateUser)

	router.GET("/users", userController.GetUsers)
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUserByID)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
