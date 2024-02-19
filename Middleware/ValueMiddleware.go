package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
func ValueMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.Query("value")
		fmt.Print(value, "valuee")
		if value == "1" {
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Value must be equal to 1"})
			c.Abort()
		}
	}
}
func BasicAuth(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, pass, ok := c.Request.BasicAuth()
		if !ok || user != username || pass != password {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}