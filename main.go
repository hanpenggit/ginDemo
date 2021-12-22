package main

import (
	"fmt"
	"ginDemo/routers"
	"ginDemo/utils"
)

func main() {
	r := routers.SetupRouter()

	var err = r.Run(fmt.Sprintf("%s%d", ":", routers.CONFIG.Port))
	if err != nil {
		utils.Logger.Error("服务启动失败,当前端口为：", err.Error())
	}
}
