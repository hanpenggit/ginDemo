package role

import "github.com/gin-gonic/gin"

func Routers(r *gin.Engine) {
	roleGrop := r.Group("role")
	roleGrop.GET("/a", aHandler)
	roleGrop.GET("/b", bHandler)
}
