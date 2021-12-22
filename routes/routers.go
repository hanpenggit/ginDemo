package routes

import (
	"encoding/json"
	"fmt"
	"ginDemo/model"
	"ginDemo/service"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Port    int    `mapstructure:"port" json:"port" ini:"port"`
	RunMode string `mapstructure:"runMode" json:"runMode" ini:"runMode" yaml:"runMode" toml:"runMode"`
	LogPath string `mapstructure:"logPath" json:"logPath" ini:"logPath" yaml:"logPath" toml:"logPath"`
}

var CONFIG = new(Config)

func IniConfigFromYaml() {
	file, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, CONFIG)
	if err != nil {
		utils.Logger.Error(err)
		return
	}
}

func SetupRouter() *gin.Engine {
	//读取配置文件
	IniConfigFromYaml()
	//初始化日志
	utils.InitLogger(CONFIG.LogPath)
	//r := gin.Default()
	r := gin.New()
	gin.SetMode(CONFIG.RunMode)
	//自定义日志，使用zap来记录日志，而不是把日志打印在控制台
	r.Use(utils.GinLogger(utils.Logger), utils.GinRecovery(utils.Logger, true))

	//简单的4种请求,返回json
	r.GET("hello", func(c *gin.Context) { service.GetHello(c) })
	r.POST("hello", func(c *gin.Context) { service.PostHello(c) })
	r.PUT("hello", func(c *gin.Context) { service.PutHello(c) })
	r.DELETE("hello", func(c *gin.Context) { service.DeleteHello(c) })

	//返回xml格式
	r.GET("/xml1", func(c *gin.Context) {
		c.XML(http.StatusOK, model.Success("Xml1", ""))
	})

	//返回xml格式
	r.GET("/xml2", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "Xml2",
			"code":    http.StatusOK,
			"data":    "",
		})
	})

	//返回yaml格式
	r.GET("/yaml1", func(c *gin.Context) {
		c.YAML(http.StatusOK, model.Success("Yaml1", ""))
	})

	//返回yaml格式
	r.GET("/yaml2", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "Yaml2",
			"code":    http.StatusOK,
			"data":    "",
		})
	})

	// getParamsFromGet?username=hanpeng&password=123123
	r.GET("getParamsFromGet", func(c *gin.Context) {

		username := c.DefaultQuery("username", "hanpeng")
		//username := c.Query("username")
		password := c.Query("password")

		data := map[string]interface{}{
			"username": username,
			"password": password,
		}

		c.JSON(http.StatusOK, model.Success("获取get请求的params参数成功", data))
	})

	//获取表单的POST请求
	r.POST("getParamsFromPost", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		username := c.DefaultPostForm("username", "123123")
		password := c.PostForm("password")
		address := c.PostForm("address")

		data := map[string]interface{}{
			"username": username,
			"password": password,
			"address":  address,
		}

		c.JSON(http.StatusOK, model.Success("获取post请求的params参数成功", data))
	})

	//读取body中的json数据
	r.POST("/json", func(c *gin.Context) {
		// 注意：下面为了举例子方便，暂时忽略了错误处理
		b, err := c.GetRawData() // 从c.Request.Body读取请求数据
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Fail("获取body参数出现异常，异常信息为："+err.Error(), ""))
		}
		// 定义map或结构体
		var m map[string]interface{}
		// 反序列化
		err = json.Unmarshal(b, &m)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Fail("格式化body参数出现异常，异常信息为："+err.Error(), ""))
		}

		c.JSON(http.StatusOK, model.Success("获取post请求的body参数成功", m))
	})

	//获取path参数   /get/hanpeng/zz
	r.GET("/get/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, model.Success("获取post请求的body参数成功", gin.H{
			"username": username,
			"address":  address,
		}))
	})

	//参数与结构体绑定
	// 绑定JSON的示例 ({"username": "hanpeng", "password": "123456"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var login model.Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, model.Success("参数绑定成功", gin.H{
				"username": login.Username,
				"password": login.Password,
			}))
		} else {
			c.JSON(http.StatusBadRequest, model.Fail("参数绑定失败", err.Error()))
		}
	})

	// 绑定form表单示例 (username=hanpeng&password=123456)
	r.POST("/loginForm", func(c *gin.Context) {
		var login model.Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, model.Success("参数绑定成功", gin.H{
				"username": login.Username,
				"password": login.Password,
			}))
		} else {
			c.JSON(http.StatusBadRequest, model.Fail("参数绑定失败", err.Error()))
		}
	})

	// 绑定QueryString示例 (/loginQuery?username=hanpeng&password=123456)
	r.GET("/loginForm", func(c *gin.Context) {
		var login model.Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, model.Success("参数绑定成功", gin.H{
				"username": login.Username,
				"password": login.Password,
			}))
		} else {
			c.JSON(http.StatusBadRequest, model.Fail("参数绑定失败", err.Error()))
		}
	})

	//单个文件上传
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Fail("上传文件失败", err.Error()))
			return
		}
		utils.Logger.Info(file.Filename)

		dst := fmt.Sprintf("D:/tmp/%s", file.Filename)
		// 上传文件到指定的目录
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Fail("保存文件失败", err.Error()))
			return
		}
		c.JSON(http.StatusOK, model.Success("文件上传成功", file.Filename))
	})

	//多个文件上传
	r.POST("/uploadTwo", func(c *gin.Context) {
		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.Fail("上传文件失败", err.Error()))
			return
		}
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("D:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			err = c.SaveUploadedFile(file, dst)
			if err != nil {
				c.JSON(http.StatusInternalServerError, model.Fail("保存文件失败", err.Error()))
				return
			}
		}
		c.JSON(http.StatusOK, model.Success("多文件上传成功", ""))
	})

	//http重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.huashuimoyu.com/")
	})

	//http转发
	r.GET("/test", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.Success("test2", ""))
	})

	//可以匹配任何方式
	r.Any("/any", func(c *gin.Context) {
		method := c.Request.Method
		c.JSON(http.StatusOK, model.Success("any调用成功，当前method:"+method, method))
	})

	//路由组,路由组也支持嵌套
	systemGroup := r.Group("/system", StatCost())

	systemGroup.GET("/a", func(c *gin.Context) { service.GetHello(c) })
	systemGroup.GET("/b", func(c *gin.Context) { service.GetHello(c) })
	systemGroup.POST("/c", func(c *gin.Context) { service.GetHello(c) })

	userGroup := systemGroup.Group("user")

	userGroup.GET("/a", func(c *gin.Context) { service.GetHello(c) })
	userGroup.GET("/b", func(c *gin.Context) { service.GetHello(c) })
	userGroup.POST("/c", func(c *gin.Context) { service.GetHello(c) })

	//自定义404的返回结果
	r.NoRoute(func(c *gin.Context) { service.PageNotFind(c) })

	return r
}

// StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "xy") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}
