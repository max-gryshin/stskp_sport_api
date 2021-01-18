package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/hetiansu5/urlquery"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/app"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/repository"
	"gitlab.com/ZmaximillianZ/stskp_sport_api/internal/routers/api"
	"net/http"
	"strconv"
	"time"
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
	appG := app.Gin{C: c}
	id, idError := strconv.Atoi(idParam)
	if idError != nil {
		appG.Response(http.StatusBadGateway, e.ERROR, idError)
	}
	user, err := repository.GetUserByID(id, models.GetAllowedUserFieldsByMethod("get"))
	if err != nil {
		appG.Response(http.StatusNotFound, e.ERROR, "resource not found")
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user)
}

// @Summary List users
// @Description Get users with params
// @Produce  json
// @Security JWT
// @Param criteria body api.QueryParams true "{"criteria": {"state": [">", "1"]},"limit": 0,"offset": 0,"order": {"id": "DESC"}}"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/users [post]
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	var queryParams api.QueryParams

	//test := api.QueryParams{
	//	[]api.CriteriaParam{
	//		{
	//			Field: "state",
	//			Value: "1",
	//			Condition: ">",
	//		},
	//	},
	//	map[string]string{"id":"DESC"},
	//	0,
	//	0,
	//}
	//testValueUrlString, _ := urlquery.Marshal(test)
	err := urlquery.Unmarshal([]byte(c.Request.URL.RawQuery), &queryParams)
	if err != nil {
		panic(err)
	}
	//if err := c.Bind(&queryParams); err != nil {
	//	panic(err)
	//}
	appG.Response(http.StatusOK, e.SUCCESS, queryParams.Offset)

	//if err := c.ShouldBindJSON(&queryParams); err != nil {
	//	logging.Error(err)
	//	appG.Response(http.StatusNotFound, e.ERROR, err)
	//	return
	//}
	//criteria, order, limit, offset, ok := api.ParseQueryParams(
	//	models.GetAllowedUserFieldsByMethod("get"),
	//	&queryParams,
	//)
	//if !ok {
	//	appG.Response(http.StatusBadGateway, e.ERROR, "invalid query params")
	//	return
	//}
	//users, err := repository.FindUserBy(criteria, order, limit, offset, models.GetAllowedUserFieldsByMethod("get"))
	//if err != nil {
	//	appG.Response(http.StatusBadGateway, e.ERROR, err)
	//	return
	//}
	//appG.Response(http.StatusOK, e.SUCCESS, users)
}

//"{\"user_name\":\"name\", \"state:1\", \"email\":\"mailname@mail.com\"}"
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
		appG.Response(http.StatusBadGateway, e.ERROR, idError)
		return
	}
	user, err := repository.GetUserByID(id, []string{})
	if err != nil {
		appG.Response(http.StatusNotFound, e.ERROR, "resource not found")
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		logging.Error(err)
		appG.Response(http.StatusNotFound, e.ERROR, err)
		return
	}
	valid := validation.Validation{}
	b, err := valid.Valid(user)
	if err != nil || !b {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, valid.Errors)
		return
	}
	if errUpdate := repository.UpdateUser(user); errUpdate != nil {
		appG.Response(http.StatusBadGateway, e.ERROR, errUpdate)
		logging.Error(errUpdate)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, user.Username)
}
