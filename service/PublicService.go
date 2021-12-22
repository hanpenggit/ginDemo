package service

import (
	"github.com/gin-gonic/gin"
)

func PageNotFind(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "The requested URL " + c.Request.RequestURI + "  (" + c.Request.Method + ") was not found on this server.",
		"data":    "",
		"code":    404,
	})
}

func GetHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Get",
	})
}
func PostHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post",
	})
}
func PutHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Put",
	})
}
func DeleteHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Delete",
	})
}
