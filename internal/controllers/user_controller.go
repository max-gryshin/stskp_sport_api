package controllers

import (
	"errors"

	"github.com/go-playground/validator"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/models"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/labstack/echo/v4"

	"net/http"
	"strconv"
	"time"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/contractions"
)

// UserController is HTTP controller for manage users
type UserController struct {
	repo         contractions.UserRepository
	errorHandler e.ErrorHandler
	BaseController
}

// NewUserController return new instance of UserController
func NewUserController(repo contractions.UserRepository, errorHandler e.ErrorHandler, v *validator.Validate) *UserController {
	return &UserController{
		repo:           repo,
		errorHandler:   errorHandler,
		BaseController: BaseController{*v},
	}
}

// GetByID return user by id
// description: Get user by id
// example: /api/v1/users/{id}/
func (ctr *UserController) GetByID(c echo.Context) error {
	var (
		id   int64
		err  error
		user models.User
	)
	if id, err = ctr.BaseController.GetID(c); err != nil {
		return err
	}
	if user, err = ctr.repo.GetByID(int(id)); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Authenticate @Summary Authenticate
// description: user authorization
// example: /api/v1/auth
func (ctr *UserController) Authenticate(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	a := models.Auth{Username: username, Password: password}
	if errValidation := ctr.BaseController.validator.Struct(&a); errValidation != nil {
		return errValidation
	}
	var (
		user  models.User
		err   error
		token string
	)
	if user, err = ctr.repo.GetByUsername(username); err != nil {
		return err
	}
	if user.InvalidPassword(password) {
		return errors.New("invalid password")
	}
	if token, err = utils.GenerateToken(a.Username, a.Password); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// GetUsers @Summary List users
// description: Get users with params
// example: /api/v1/users
func (ctr *UserController) GetUsers(c echo.Context) error {
	var (
		users models.Users
		err   error
	)
	if users, err = ctr.repo.GetUsers(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// Create @Summary Create user
// description: Create user
// example: /api/v1/create
func (ctr *UserController) Create(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")
	a := models.Auth{Username: username, Password: password}
	if errValidation := ctr.BaseController.validator.Struct(&a); errValidation != nil {
		return errValidation
	}
	var (
		userExist models.User
		err       error
		user      models.User
	)
	if userExist, err = ctr.repo.GetByUsername(a.Username); err != nil {
		return err
	}
	if userExist.ID != 0 {
		return errors.New("user with username " + a.Username + " exists")
	}
	user = models.User{Username: a.Username, State: models.StateHalfRegistration, CreatedAt: time.Now()}
	if errSetPassword := user.SetPassword(a.Password); errSetPassword != nil {
		return errSetPassword
	}
	if errCreateUser := ctr.repo.CreateUser(&user); errCreateUser != nil {
		return errCreateUser
	}
	return c.JSON(http.StatusOK, map[string]string{"id": strconv.Itoa(user.ID), "username": user.Username})
}

// Update return user by id
// description: Update user by id
// example: /api/v1/users/{id}/
func (ctr *UserController) Update(c echo.Context) error {
	var (
		id   int64
		err  error
		user models.User
	)
	if id, err = ctr.BaseController.GetID(c); err != nil {
		return err
	}
	if user, err = ctr.repo.GetByID(int(id)); err != nil {
		return err
	}
	if errBinding := c.Bind(&user); errBinding != nil {
		return errBinding
	}
	if errValidate := ctr.BaseController.validator.Struct(user); errValidate != nil {
		return errValidate
	}
	if errUpdateUser := ctr.repo.UpdateUser(&user); errUpdateUser != nil {
		return errUpdateUser
	}
	return c.JSON(http.StatusOK, user)
}

// Delete return user by id
// description: Delete user by id
// example: /api/v1/users/{id}/ [delete]
func (ctr *UserController) Delete(c echo.Context) error {
	var (
		id   int64
		err  error
		user models.User
	)
	if id, err = ctr.BaseController.GetID(c); err != nil {
		return err
	}
	if user, err = ctr.repo.GetByID(int(id)); err != nil {
		return err
	}
	if errDelete := ctr.repo.Delete(&user); errDelete != nil {
		return errDelete
	}
	return c.JSON(http.StatusOK, "OK")
}
