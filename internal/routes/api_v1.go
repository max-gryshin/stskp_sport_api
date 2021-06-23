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
	router.POST("/create", userController.Create)
	user := router.Group("/users")
	user.GET("/", userController.GetUsers, jwt)
	user.POST("/", userController.Create, jwt)
	user.GET("/:id", userController.GetByID)
	user.PUT("/:id", userController.Update)
	user.DELETE("/:id", userController.Delete, jwt)

	workout := router.Group("/workout")

	workout.Group("/")
}
