package routers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/db"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/middleware/jwt"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	v1 "github.com/ZmaximillianZ/stskp_sport_api/internal/routers/api/v1"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	userRepo := repository.NewUserRepository(&db.DB, goqu.Dialect("postgres"))
	userController := controllers.NewUserController(userRepo)

	//url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	apiv1.GET("/users", v1.GetUsers)
	apiv1.GET("users/:id", userController.GetUserByID)
	apiv1.PATCH("/users/:id/update", v1.UpdateUser)

	return router
}
