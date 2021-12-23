package role

import (
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	roleGrop := r.Group("role")
	roleGrop.Use(utils.JWTAuthMiddleware())
	roleGrop.GET("/a", aHandler)
	roleGrop.GET("/b", bHandler)
}
