package api

import (
	"net/http"

	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/repository"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/app"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/util"
)

type Auth struct {
	Username string `json:"username" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}

// @Summary Get Auth
// @Description user authorization
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Header 200 {string} Access-Control-Allow_origin "*"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/user/auth [post]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c} // TODO: Do not wrap git contex. use method directly
	valid := validation.Validation{}

	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")
	a := Auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	if !ok {
		if valid.HasErrors() {
			// maybe c.Error(valid.Errors)
			app.MarkErrors(valid.Errors)
		}
		// FIXME c c.AbortWithStatus(http.StatusBadRequest)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	user, err := repository.FindUserByUsername(username)
	if err != nil {
		logging.Error(err)
	}
	invalid := user.InvalidPassword(password)

	if invalid {
		// FIXME c c.AbortWithStatus(http.StatusBadRequest)
		appG.Response(http.StatusBadRequest, e.ERROR_AUTH_CHECK_CREDENTIALS_FAIL, nil)
		return
	}

	token, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		// FIXME c c.AbortWithStatus(http.StatusInternalServerError)
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	// FIXME: This looks better:
	// c.JSON(e.SUCCESS, map[string]string{
	// 	"token": token,
	// })
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
