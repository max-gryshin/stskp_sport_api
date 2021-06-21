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
	router.GET("/users/", userController.GetUsers, jwt) // jwt do not work
	router.POST("/users/", userController.CreateUser, jwt)
	router.GET("/users/:id", userController.GetUserByID, jwt)
	router.PUT("/users/:id", userController.UpdateUser)
	router.DELETE("/users/:id", userController.DeleteUser, jwt)
}
