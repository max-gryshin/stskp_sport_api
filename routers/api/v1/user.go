package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/models"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/app"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/e"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/logging"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/repository"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/routers/api"
	"net/http"
	"time"
)

func GetUser(c *gin.Context) {
	appG := app.Gin{C: c}
	article := "Username"
	appG.Response(http.StatusOK, e.SUCCESS, article)
}

func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")
	a := api.Auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	if !ok {
		if valid.HasErrors() {
			app.MarkErrors(valid.Errors)
		}
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	user := models.User{Username: a.Username, State: models.StateHalfRegistration, CreatedAt: time.Now()}
	if err := user.SetPassword(a.Password); err != nil {
		logging.Error(err)
	}
	if errSave := repository.CreateUser(user); errSave != nil {
		appG.Response(http.StatusBadGateway, e.ERROR, errSave)
		logging.Error(errSave)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user.Username)
}
