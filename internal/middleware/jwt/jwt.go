package jwt

import (
	"net/http"

	"github.com/ZmaximillianZ/stskp_sport_api/internal/e"
	"github.com/ZmaximillianZ/stskp_sport_api/internal/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.Success
		// TODO: recommended use Authentication: Bearer
		// https://swagger.io/docs/specification/authentication/bearer-authentication/
		token := c.GetHeader("X-AUTH-TOKEN")
		if token == "" {
			code = e.InvalidParams
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = e.ErrorAuthCheckTokenTimeout
				default:
					code = e.ErrorAuthCheckTokenFail
				}
			}
		}

		if code != e.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
