package routes

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterAuth initialize routing information
func RegisterAuth(router gin.IRoutes, userController *controllers.UserController) {
	router.POST("/api/user/auth", userController.Authenticate)
	router.POST("/api/user/create", userController.CreateUser)
}
