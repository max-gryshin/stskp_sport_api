package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/app"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/e"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/pkg/util"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/service/auth_service"
)

type auth struct {
	Username string `json:"username" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		if valid.HasErrors() {
			app.MarkErrors(valid.Errors)
		}
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: a.Username, Password: a.Password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, err.Error())
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
