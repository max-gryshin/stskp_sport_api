package v1

import (
	"net/http"
	"strconv"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/app"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/routers/api"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//{
//    \"criteria\": {
//        \"state\": [\">\", \"1\"]
//    },
//    \"limit\": 1,
//    \"offset\": 0,
//    \"order\": {\"id\": \"DESC\"}
//}
// @Summary List users
// @Description Get users with params
// @Produce  json
// @Security JWT
// @Success 200 {array} models.User
// @Failure 500 {object} app.Response
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	var queryParams api.QueryParams
	if err := c.ShouldBindJSON(&queryParams); err != nil {
		logging.Error(err)
		appG.Response(http.StatusNotFound, e.Error, err)
		return
	}
	// TODO: code after commit is was be repeated very often try export it to middleware
	criteria, order, limit, offset, ok := api.ParseQueryParams(
		models.GetAllowedUserFieldsByMethod("get"),
		&queryParams,
	)
	if !ok {
		appG.Response(http.StatusBadGateway, e.Error, "invalid query params")
		return
	}
	users, err := repository.FindUserBy(criteria, order, limit, offset, models.GetAllowedUserFieldsByMethod("get"))
	if err != nil {
		appG.Response(http.StatusBadGateway, e.Error, err)
		return
	}
	appG.Response(http.StatusOK, e.Success, users)
}

// @Summary Update user
// @Description Update user
// @Accept  json
// @Produce  json
// @Security JWT
// @Param id path int true "User ID"
// @Param user body models.User true "update user_name, state, email"
// @Success 200 {object} models.User
// @Header 200 {string} X-AUTH-TOKEN "qwerty"
// @Failure 500 {object} app.Response
// @Router /api/v1/users/{id}/update [patch]
func UpdateUser(c *gin.Context) {
	appG := app.Gin{C: c}
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		appG.Response(http.StatusBadGateway, e.Error, idError)
		return
	}
	user, err := repository.GetUserByID(id, []string{})
	if err != nil {
		appG.Response(http.StatusNotFound, e.Error, "resource not found")
		return
	}
	if errBindingUserToJSON := c.ShouldBindJSON(&user); errBindingUserToJSON != nil {
		logging.Error(errBindingUserToJSON)
		appG.Response(http.StatusNotFound, e.Error, errBindingUserToJSON)
		return
	}
	valid := validation.Validation{}
	b, err := valid.Valid(user)
	if err != nil || !b {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.InvalidParams, valid.Errors)
		return
	}
	if errUpdate := repository.UpdateUser(&user); errUpdate != nil {
		appG.Response(http.StatusBadGateway, e.Error, errUpdate)
		logging.Error(errUpdate)
		return
	}
	appG.Response(http.StatusOK, e.Success, user.Username)
}
