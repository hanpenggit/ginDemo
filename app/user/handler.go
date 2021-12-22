package user

import (
	"ginDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func aHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("user_a", ""))
}

func bHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("user_b", ""))
}
