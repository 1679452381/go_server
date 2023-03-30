/**
 * @author X
 * @date 2023/3/30
 * @description 入口文件
 **/
package main

import (
	"fmt"
	"go_server/core"
	"go_server/core/zlog"
	"go_server/global"
	"go_server/routers"
)

func main() {
	//配置文件初始化
	core.InitConf()
	//连接数据库
	global.DB = core.InitGrom()
	//	日志文件初始化
	zlog.Init()
	//	路由初始化
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	fmt.Println(addr)
	router.Run(addr)
}
