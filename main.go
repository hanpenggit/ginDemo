package main

import (
	"fmt"
	"ginDemo/routes"
	"ginDemo/utils"
)

func main() {
	r := routes.SetupRouter()

	var err = r.Run(fmt.Sprintf("%s%d", ":", routes.CONFIG.Port))
	if err != nil {
		utils.Logger.Error("服务启动失败,当前端口为：", err.Error())
	}
}
