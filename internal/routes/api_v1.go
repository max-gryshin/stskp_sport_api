package routes

import (
	. "github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// RegisterAPIV1 initialize routing information
func RegisterAPIV1(router gin.IRouter, userController *UserController) {
	router.Use(jwt.JWT())

	users := router.Group("users")

	users.GET("/", userController.GetUsers)
	users.POST("/", userController.CreateUser)
	users.GET("/:id", userController.GetUserByID)
	users.PUT("/:id", userController.UpdateUser)
	users.DELETE("/:id", userController.DeleteUser)
}
