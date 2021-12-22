package role

import (
	"ginDemo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func aHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("role_a", ""))
}

func bHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("role_b", ""))
}
