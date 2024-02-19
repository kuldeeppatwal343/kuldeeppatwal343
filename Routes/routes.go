package routes

import (
	middleware "Gol/Middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)
func Get(){
	r:=  gin.Default()
	r.Use(middleware.BasicAuth("admin", "password"))
	r.GET("/get", func(c *gin.Context) {
		value := c.Query("value")
	c.JSON(http.StatusOK, gin.H{"value": value})
	})
	r.Run(":8080")
}