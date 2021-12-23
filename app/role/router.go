package role

import (
	"fmt"
	"ginDemo/routers"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) {
	roleGrop := r.Group("role")
	fmt.Println(routers.JwtAuth)
	//添加权限验证
	roleGrop.Use(routers.JwtAuth.MiddlewareFunc())
	roleGrop.GET("/a", aHandler)
	roleGrop.GET("/b", bHandler)
}
