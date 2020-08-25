package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/app"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/e"
	"net/http"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	article := "Username"
	appG.Response(http.StatusOK, e.SUCCESS, article)
}
