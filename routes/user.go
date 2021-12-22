package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadUser(r *gin.Engine) {
	roleGroup := r.Group("user")
	roleGroup.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "a",
			"code":    http.StatusOK,
			"data":    "",
		})
	})
	roleGroup.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
			"code":    http.StatusOK,
			"data":    "",
		})
	})
	roleGroup.GET("/c", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "c",
			"code":    http.StatusOK,
			"data":    "",
		})
	})
}
