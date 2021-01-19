package api

import (
	"net/http"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/app"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
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
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err := repository.FindUserByUsername(username)
	if err != nil {
		logging.Error(err)
	}
	invalid := user.InvalidPassword(password)

	if invalid {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := util.GenerateToken(a.Username, a.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(e.SUCCESS, map[string]string{"token": token})
}
