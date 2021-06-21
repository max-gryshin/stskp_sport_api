package controllers

import (
	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/logging"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/astaxie/beego/validation"
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/app"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/contractions"
)

// UserController is HTTP controller for manage users
type UserController struct {
	repo         contractions.UserRepository
	validation   validation.Validation
	errorHandler e.ErrorHandler
}

// NewUserController return new instance of UserController
func NewUserController(repo contractions.UserRepository, v validation.Validation, errorHandler e.ErrorHandler) *UserController {
	return &UserController{repo, v, errorHandler}
}

// GetUserByID return user by id
// description: Get user by id
// example: /api/v1/users/{id}/ [get]
func (ctr *UserController) GetUserByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := ctr.repo.GetByID(int(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// Authenticate @Summary Authenticate
// description: user authorization
// example: /api/v1/auth [post]
func (ctr *UserController) Authenticate(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	a := models.Auth{Username: username, Password: password}
	ok, _ := ctr.validation.Valid(&a)

	if !ok {
		var errors string
		if ctr.validation.HasErrors() {
			// maybe e.Error(valid.Errors)
			errors = app.MarkErrors(ctr.validation.Errors, true)
		}
		return c.String(http.StatusBadRequest, errors)
	}
	user, err := ctr.repo.GetByUsername(username)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	invalid := user.InvalidPassword(password)

	if invalid {
		return c.String(http.StatusBadRequest, "invalid password")
	}

	token, err := utils.GenerateToken(a.Username, a.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// GetUsers @Summary List users
// description: Get users with params
// example: /api/v1/users [get]
func (ctr *UserController) GetUsers(c echo.Context) error {
	users, err := ctr.repo.GetUsers()
	if err != nil {
		logging.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

// CreateUser @Summary Create user
// description: Create user
// example: /api/v1/create [post]
func (ctr *UserController) CreateUser(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	a := models.Auth{Username: username, Password: password}
	ok, _ := ctr.validation.Valid(&a)
	if !ok {
		var errors string
		if ctr.validation.HasErrors() {
			errors = app.MarkErrors(ctr.validation.Errors, true)
		}
		return c.String(http.StatusBadRequest, errors)
	}
	userExist, err := ctr.repo.GetByUsername(a.Username)
	if err != nil {
		logging.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if userExist.ID != 0 {
		return c.JSON(http.StatusConflict, map[string]string{"message": "user exists", "details": a.Username})
	}
	user := models.User{Username: a.Username, State: models.StateHalfRegistration, CreatedAt: time.Now()}
	if err := user.SetPassword(a.Password); err != nil {
		logging.Error(err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	if errSave := ctr.repo.CreateUser(&user); errSave != nil {
		logging.Error(errSave)
		return ctr.errorHandler.Handle(c, errSave)
	}

	return c.JSON(http.StatusOK, map[string]string{"id": strconv.Itoa(user.ID), "username": user.Username})
}

// UpdateUser return user by id
// description: Update user by id
// example: /api/v1/users/{id}/ [put]
func (ctr *UserController) UpdateUser(c echo.Context) error {
	id, idError := strconv.Atoi(c.Param("id"))
	if idError != nil {
		return c.String(http.StatusInternalServerError, idError.Error())
	}
	user, err := ctr.repo.GetByID(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if errBindingUserToJSON := c.Bind(&user); errBindingUserToJSON != nil {
		logging.Error(errBindingUserToJSON)
		return c.String(http.StatusInternalServerError, errBindingUserToJSON.Error())
	}
	b, err := ctr.validation.Valid(user)
	var errorResults string
	if err != nil || !b {
		errorResults = app.MarkErrors(ctr.validation.Errors, true)
		return c.String(http.StatusInternalServerError, errorResults)
	}
	if errUpdate := ctr.repo.UpdateUser(&user); errUpdate != nil {
		logging.Error(errUpdate)
		return c.String(http.StatusInternalServerError, errUpdate.Error())
	}
	return c.JSON(http.StatusOK, user)
}

// DeleteUser return user by id
// description: Delete user by id
// example: /api/v1/users/{id}/ [delete]
func (ctr *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := ctr.repo.GetByID(int(id))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if errDelete := ctr.repo.DeleteUser(&user); errDelete != nil {
		return c.String(http.StatusInternalServerError, errDelete.Error())
	}
	return c.JSON(http.StatusOK, "OK")
}
