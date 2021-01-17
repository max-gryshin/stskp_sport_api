package cors

import (
	"github.com/gin-gonic/gin"
)

// CORS is cors middleware
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // TODO: possible only for debug
		c.Header("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		c.Next()
	}
}
