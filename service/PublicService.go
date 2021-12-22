package service

import (
	"ginDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PageNotFind(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "The requested URL " + c.Request.RequestURI + "  (" + c.Request.Method + ") was not found on this server.",
		"data":    "",
		"code":    http.StatusNotFound,
	})
}

func GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("Get", ""))
}
func PostHello(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("Post", ""))
}
func PutHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Put",
		"code":    http.StatusOK,
		"data":    "",
	})
}
func DeleteHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete",
		"code":    http.StatusOK,
		"data":    "",
	})
}
