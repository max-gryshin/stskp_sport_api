package routers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/db"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/middleware/jwt"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

// RegisterAPIV1 initialize routing information
func RegisterAPIV1(router gin.IRoutes) {
	router.Use(jwt.JWT())
	queryBuilder := goqu.Dialect("postgres")
	userRepo := repository.NewUserRepository(&db.DB, queryBuilder)
	userController := controllers.NewUserController(userRepo)

	router.GET("/users", userController.GetUsers)
	router.POST("/users", userController.CreateUser)
	router.GET("/users/:id", userController.GetUserByID)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser)
}
