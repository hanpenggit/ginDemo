package utils

import (
	"ginDemo/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strings"
	"time"
)

const TokenExpireDuration = time.Hour * 1

var MySecret = []byte("myjwtsecret")

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username   string   `json:"username"`
	Role       []string `json:role`
	Permission []string `json:permission`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(username string, role []string, permission []string) (string, error, int64) {
	expire := time.Now().Add(TokenExpireDuration).Unix()
	// 创建一个我们自己的声明
	c := MyClaims{
		username,   // 自定义字段
		role,       // 自定义字段
		permission, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: expire,    // 过期时间
			Issuer:    "ginDemo", // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	tokenValue, err := token.SignedString(MySecret)
	return tokenValue, err, expire
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
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
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.Success("无效的token", ""))
			c.Abort()
			return
		}
		var uri = c.Request.RequestURI
		var method = c.Request.Method
		for _, j := range mc.Permission {
			if uri == j {
				// 校验通过，存放用户信息
				c.Set("username", mc.Username)
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, model.Success("您没有该权限: "+uri+"   ("+method+")", ""))
		c.Abort()
	}
}
