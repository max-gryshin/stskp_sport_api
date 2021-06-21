package routes

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RegisterAPIV1 initialize routing information
func RegisterAPIV1(router *echo.Group, userController *controllers.UserController) {
	jwt := middleware.JWT([]byte("get_key_from_env"))

	router.POST("/auth", userController.Authenticate)
	router.POST("/create", userController.CreateUser)
	user := router.Group("/users")
	user.GET("/", userController.GetUsers, jwt) // jwt do not work
	user.POST("/", userController.CreateUser, jwt)
	user.GET("/:id", userController.GetUserByID, jwt)
	user.PUT("/:id", userController.UpdateUser)
	user.DELETE("/:id", userController.DeleteUser, jwt)
}
