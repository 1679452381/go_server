/**
 * @author X
 * @date 2023/3/30
 * @description 入口文件
 **/
package main

import (
	"go_server/core"
	logger "go_server/core/zlog"
	"go_server/global"
)

func main() {
	//配置文件初始化
	core.InitConf()
	//连接数据库
	global.DB = core.InitGrom()
	//	日志文件初始化
	logger.Init()

}
