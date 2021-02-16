package routes

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// RegisterAPIV1 initialize routing information
func RegisterAPIV1(router gin.IRouter, userController *controllers.UserController) {
	users := router.Group("users")
	router.POST("/auth", userController.Authenticate)
	router.POST("/create", userController.CreateUser)
	users.GET("/", userController.GetUsers).Use(jwt.JWT())
	users.POST("/", userController.CreateUser).Use(jwt.JWT())
	users.GET("/:id", userController.GetUserByID).Use(jwt.JWT())
	users.PUT("/:id", userController.UpdateUser).Use(jwt.JWT())
	users.DELETE("/:id", userController.DeleteUser).Use(jwt.JWT())
}
