/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package routers

import (
	"github.com/gin-gonic/gin"
	"go_server/global"
)

//
//type RouterGroup struct {
//	*gin.RouterGroup
//}

type ApiRouter struct {
}

var RouterGroupApp = new(ApiRouter)

func InitRouter() *gin.Engine {
	//Gin 框架提供了三种运行模式，分别是 "debug"、"release" 和 "test"
	//在 "debug" 模式下，Gin 框架会打印更详细的日志和错误信息，适用于开发和调试阶段。
	//在 "release" 模式下，Gin 框架会尽量减少日志和错误信息的输出，以提高性能，适用于生产环境。
	//在 "test" 模式下，Gin 框架会关闭日志输出，以方便进行测
	gin.SetMode(global.Config.System.Env)

	router := gin.Default()
	apiRouterGroup := router.Group("api")
	//settingsApi := api.ApiGroupApp.SettingsApi
	//
	//router.GET("/", settingsApi.SettingsInfoView)

	//routerGroup := RouterGroup{router}
	//routerGroup.SettingRouter()

	RouterGroupApp.SiteInfoRouter(apiRouterGroup)

	return router
}
