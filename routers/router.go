package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/middleware/jwt"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/routers/api"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/auth", api.GetAuth)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/user", v1.GetUser)
	}

	return router
}
