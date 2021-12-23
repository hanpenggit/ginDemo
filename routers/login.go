package routers

import (
	"ginDemo/model"
	"ginDemo/utils"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

//https://github.com/appleboy/gin-jwt
//jwt密钥
var identityKey = "hanpeng_key"

func InitJwt(r *gin.Engine) *jwt.GinJWTMiddleware {
	// the jwt middleware
	duthMiddleware, jwtErr := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals model.Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			UserName := loginVals.UserName
			password := loginVals.Password

			if (UserName == "admin" && password == "admin") || (UserName == "test" && password == "test") {
				return &model.User{
					UserName: UserName,
					NickName: UserName + "  Ni",
					Remarks:  "Hi " + UserName,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if jwtErr != nil {
		utils.Logger.Error("jwt.New() Error:" + jwtErr.Error())
	}

	errInit := duthMiddleware.MiddlewareInit()

	if errInit != nil {
		utils.Logger.Error("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/login", duthMiddleware.LoginHandler)

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", duthMiddleware.RefreshHandler)
	return duthMiddleware
}
