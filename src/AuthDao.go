package src

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, status := c.GetQuery("token"); !status {
			c.String(http.StatusUnauthorized, "You are not authorized to  operate")
			c.Abort()
		}
	}
}
