package user

import (
	"ginDemo/model"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func aHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("user_a", ""))
}

func bHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.Success("user_b", ""))
}

func loginHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Fail("输入的参数有误", err.Error()))
		return
	}
	// 校验用户名和密码是否正确
	if user.Username == "admin" && user.Password == "admin" {
		// 生成Token
		tokenString, _ := utils.GenToken(user.Username)
		c.JSON(http.StatusOK, model.Success("登录成功", gin.H{"token": tokenString}))
		return
	}
	c.JSON(http.StatusOK, model.Success("账号或密码错误", ""))
	return
}

func userDetailHandler(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, model.Success("缺少认证信息", ""))
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusUnauthorized, model.Success("token格式不正确", ""))
		c.Abort()
		return
	}
	// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
	mc, err := utils.ParseToken(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Success("无效的token", ""))
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, model.Success("获取用户信息成功", mc))
	return
}
