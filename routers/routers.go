package routers

import (
	"ginDemo/service"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
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

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	//读取配置文件
	IniConfigFromYaml()
	//初始化日志
	utils.InitLogger(CONFIG.LogPath)
	//r := gin.Default()
	r := gin.New()
	gin.SetMode(CONFIG.RunMode)
	//自定义日志，使用zap来记录日志，而不是把日志打印在控制台
	r.Use(utils.GinLogger(utils.Logger), utils.GinRecovery(utils.Logger, true))
	for _, opt := range options {
		opt(r)
	}
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
