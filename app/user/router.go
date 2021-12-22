package user

import "github.com/gin-gonic/gin"

func Routers(r *gin.Engine) {
	userGroup := r.Group("/user")
	userGroup.GET("/a", aHandler)
	userGroup.GET("/b", bHandler)
}
