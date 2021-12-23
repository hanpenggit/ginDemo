package user

import (
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	//登录
	r.POST("/login", loginHandler)

	userGroup := r.Group("/user")
	//获取用户信息
	userGroup.POST("/detail", userDetailHandler)
	userGroup.GET("/a", aHandler)
	userGroup.GET("/b", bHandler)
}
