package routers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/db"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/middleware/jwt"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/routers/api"
	v1 "github.com/ZmaximillianZ/stskp_sport_api/internal/routers/api/v1"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter(conf *setting.Setting) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	connection := db.CreateConnection(db.ConnectionSettions{
		URL:         conf.DBConfig.URL,
		MaxIdleCons: 100,
		MaxOpenCons: 10,
	})

	queryBuilder := goqu.Dialect("postgres")

	router.POST("/api/user/auth", api.GetAuth)
	router.POST("/api/user/create", v1.CreateUser)
	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())

	userRepo := repository.NewUserRepository(connection, queryBuilder)
	userController := controllers.NewUserController(userRepo)
	apiv1.GET("/users", userController.GetUsers)
	apiv1.GET("users/:id", userController.GetUserByID)
	apiv1.PATCH("/users/:id/update", v1.UpdateUser)

	return router
}
