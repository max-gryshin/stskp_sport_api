package e

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
)

type ErrorHandler struct {
}

func (eh *ErrorHandler) Handle(c *gin.Context, e error) {
	if pg := handlePgError(e); pg != nil {
		c.JSON(
			InvalidParams,
			map[string]string{
				"detail":  pg.Detail,
				"message": pg.Message,
				"code":    pg.Code,
			},
		)
		return
	}
}

func handlePgError(e error) *pgconn.PgError {
	if err, ok := e.(*pgconn.PgError); ok {
		return err
	}
	return nil
}
