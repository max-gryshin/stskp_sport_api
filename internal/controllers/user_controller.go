package controllers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/astaxie/beego/validation"

	"net/http"
	"strconv"
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/app"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/contractions"
	"github.com/gin-gonic/gin"
)

// UserController is HTTP controller for manage users
type UserController struct {
	repo       contractions.UserRepository
	validation validation.Validation
}

// NewUserController return new instance of UserController
func NewUserController(repo contractions.UserRepository, v validation.Validation) *UserController {
	return &UserController{repo, v}
}

// GetUserByID return user by id
// @Summary Show a user
// @Description Get user by id
// @Produce  json
// @Security JWT
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Header 200 {string} X-AUTH-TOKEN "qwerty"
// @Failure 500 {object} app.Response
// @Router /api/v1/users/{id}/ [get]
func (ctr *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := ctr.repo.GetByID(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Authenticate
// @Description user authorization
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Header 200 {string} Access-Control-Allow_origin "*"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/user/auth [post]
func (ctr *UserController) Authenticate(c *gin.Context) {
	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")
	a := models.Auth{Username: username, Password: password}
	ok, _ := ctr.validation.Valid(&a)

	if !ok {
		if ctr.validation.HasErrors() {
			// maybe c.Error(valid.Errors)
			app.MarkErrors(ctr.validation.Errors)
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	user, err := ctr.repo.GetByUsername(username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	invalid := user.InvalidPassword(password)

	if invalid {
		c.AbortWithStatusJSON(http.StatusBadRequest, "invalid password")
		return
	}

	token, err := utils.GenerateToken(a.Username, a.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"token": token})
}

// @Summary List users
// @Description Get users with params
// @Produce  json
// @Security JWT
// @Success 200 {array} models.User
// @Failure 500 {object} app.Response
// @Router /api/v1/users [get]
func (ctr *UserController) GetUsers(c *gin.Context) {
	users, err := ctr.repo.GetUsers()
	if err != nil {
		logging.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Create user
// @Description Create user
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/user/create [post]
func (ctr *UserController) CreateUser(c *gin.Context) {
	username, _ := c.GetQuery("username")
	password, _ := c.GetQuery("password")
	a := models.Auth{Username: username, Password: password}
	ok, _ := ctr.validation.Valid(&a)
	if !ok {
		if ctr.validation.HasErrors() {
			app.MarkErrors(ctr.validation.Errors)
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
	if errSave := ctr.repo.CreateUser(&user); errSave != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		logging.Error(errSave)
		return
	}
	c.JSON(e.Success, map[string]string{"id": strconv.Itoa(user.ID), "username": user.Username})
}

func (ctr *UserController) UpdateUser(c *gin.Context) {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		c.AbortWithStatus(http.StatusInternalServerError)

		return
	}
	user, err := ctr.repo.GetByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if errBindingUserToJSON := c.ShouldBindJSON(&user); errBindingUserToJSON != nil {
		logging.Error(errBindingUserToJSON)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	b, err := ctr.validation.Valid(user)
	if err != nil || !b {
		app.MarkErrors(ctr.validation.Errors)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if errUpdate := ctr.repo.UpdateUser(&user); errUpdate != nil {
		logging.Error(errUpdate)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ctr *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := ctr.repo.GetByID(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	if errDelete := ctr.repo.DeleteUser(&user); errDelete != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, "OK")
}
