package e

import (
	"net/http"

	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
)

type ErrorHandler struct {
}

func (eh *ErrorHandler) Handle(c echo.Context, e error) error {
	if pg := handlePgError(e); pg != nil {
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{
				"detail":  pg.Detail,
				"message": pg.Message,
				"code":    pg.Code,
			},
		)
	}
	return nil
}

func handlePgError(e error) *pgconn.PgError {
	if err, ok := e.(*pgconn.PgError); ok {
		return err
	}
	return nil
}
