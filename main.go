package main

import (
	"fmt"
	"ginDemo/service"
	"ginDemo/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Port    int    `mapstructure:"port" json:"port" ini:"port"`
	RunMode string `mapstructure:"runMode" json:"runMode" ini:"runMode" yaml:"runMode" toml:"runMode"`
	LogPath string `mapstructure:"logPath" json:"logPath" ini:"logPath" yaml:"logPath" toml:"logPath"`
}

var (
	CONFIG = new(Config)
)

func main() {
	IniConfigFromYaml()
	utils.InitLogger(CONFIG.LogPath)
	//r := gin.Default()
	r := gin.New()
	gin.SetMode(CONFIG.RunMode)
	//自定义日志，使用zap来记录日志，而不是把日志打印在控制台
	r.Use(utils.GinLogger(utils.Logger), utils.GinRecovery(utils.Logger, true))

	//简单的4种请求
	r.GET("hello", func(c *gin.Context) { service.GetHello(c) })
	r.POST("hello", func(c *gin.Context) { service.PostHello(c) })
	r.PUT("hello", func(c *gin.Context) { service.PutHello(c) })
	r.DELETE("hello", func(c *gin.Context) { service.DeleteHello(c) })

	//自定义404的返回结果
	r.NoRoute(func(c *gin.Context) { service.PageNotFind(c) })

	var err = r.Run(fmt.Sprintf("%s%d", ":", CONFIG.Port))
	if err != nil {
		utils.Logger.Error("服务启动失败,当前端口为：", err.Error())
	}
}

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
