package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/app"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/routers/api"
	"net/http"
	"time"
)

func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	users, err := repository.UserAll()
	if err != nil {
		appG.Response(http.StatusBadGateway, e.ERROR, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, users)
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
