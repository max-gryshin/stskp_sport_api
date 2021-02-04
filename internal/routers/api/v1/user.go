package v1

import (
	"github.com/fatih/structs"

	"net/http"
	"strconv"
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/app"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/routers/api"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// @Summary Create user
// @Description Create user
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/user/create [post]
func CreateUser(c *gin.Context) {
	valid := validation.Validation{}
	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")
	a := api.Auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)
	if !ok {
		if valid.HasErrors() {
			app.MarkErrors(valid.Errors)
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user := models.User{Username: a.Username, State: models.StateHalfRegistration, CreatedAt: time.Now()}
	if err := user.SetPassword(a.Password); err != nil {
		logging.Error(err)
		c.AbortWithStatus(http.StatusBadRequest) // OR: use c.AbortWithError()
		return
	}
	if errSave := repository.CreateUser(&user); errSave != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		logging.Error(errSave)
		return
	}
	c.JSON(e.Success, map[string]string{"id": strconv.Itoa(user.ID), "username": user.Username})
}

// @Summary Show a user
// @Description Get user by id
// @Produce  json
// @Security JWT
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Header 200 {string} X-AUTH-TOKEN "qwerty"
// @Failure 500 {object} app.Response
// @Router /api/v1/users/{id}/ [get]
func GetUser(c *gin.Context) {
	idParam := c.Param("id")
	id, idError := strconv.Atoi(idParam)
	if idError != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err := repository.GetUserByID(id, models.GetAllowedUserFieldsByMethod("get"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, structs.Map(user)) // FIXME: may be just struct
}

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
