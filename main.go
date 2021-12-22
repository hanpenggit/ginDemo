package main

import (
	"fmt"
	"ginDemo/app/role"
	"ginDemo/app/user"
	"ginDemo/routers"
	"ginDemo/utils"
)

func main() {
	routers.Include(role.Routers, user.Routers)

	// 初始化路由
	r := routers.Init()

	var err = r.Run(fmt.Sprintf("%s%d", ":", routers.CONFIG.Port))

	if err != nil {
		utils.Logger.Error("服务启动失败,当前端口为：", err.Error())
	}
}
