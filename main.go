/**
 * @author X
 * @date 2023/3/30
 * @description 入口文件
 **/
package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_server/core"
	"go_server/core/zlog"
	_ "go_server/docs"
	"go_server/global"
	"go_server/routers"
)

// @title gvb_gin_server
// @version 1.0
// @description  API文档
// @host   127.0.0.1:8080
// @BasePath /
func main() {
	//配置文件初始化
	core.InitConf()
	//连接数据库
	global.DB = core.InitGrom()
	//	日志文件初始化
	zlog.Init()
	//	路由初始化
	router := routers.InitRouter()
	// swagger 初始化
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(global.Config.System.Addr())
}
