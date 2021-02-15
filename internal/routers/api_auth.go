package routers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/controllers"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/db"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/setting"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
)

// RegisterAuth initialize routing information
func RegisterAuth(router gin.IRoutes) {
	queryBuilder := goqu.Dialect("postgres")
	userRepo := repository.NewUserRepository(db.CreateDBConnection(&setting.AppSetting.DBConfig), queryBuilder)
	userController := controllers.NewUserController(userRepo)
	router.POST("/api/user/auth", userController.GetAuth)
	router.POST("/api/user/create", userController.CreateUser)
}
