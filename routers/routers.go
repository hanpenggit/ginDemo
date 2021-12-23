package routers

import (
	"ginDemo/service"
	"ginDemo/utils"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

var JwtAuth *jwt.GinJWTMiddleware

// 初始化
func Init() *gin.Engine {
	//读取配置文件
	IniConfigFromYaml()
	//初始化日志
	utils.InitLogger(CONFIG.LogPath)
	//r := gin.Default()
	r := gin.New()
	gin.SetMode(CONFIG.RunMode)

	//加载Jwt对象
	JwtAuth = InitJwt(r)

	//自定义日志，使用zap来记录日志，而不是把日志打印在控制台
	r.Use(utils.GinLogger(utils.Logger), utils.GinRecovery(utils.Logger, true))
	for _, opt := range options {
		opt(r)
	}
	//自定义404的返回结果
	r.NoRoute(func(c *gin.Context) { service.PageNotFind(c) })
	return r
}
