package controllers

import (
	"errors"
	"strconv"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
)

type BaseController struct {
	validator validator.Validate
}

func (ctr *BaseController) GetID(c echo.Context) (int64, error) {
	var (
		id  int64
		err error
	)
	if id, err = strconv.ParseInt(c.Param("id"), 10, 64); err != nil {
		return id, errors.New("invalid id")
	}

	return id, nil
}

func (ctr *BaseController) BindAndValidate(c echo.Context, model interface{}) error {
	if errBinding := c.Bind(&model); errBinding != nil {
		return errBinding
	}
	if errValidate := ctr.validator.Struct(model); errValidate != nil { //todo: fix for workout
		return errValidate
	}

	return nil
}
