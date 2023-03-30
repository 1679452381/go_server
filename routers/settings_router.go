/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package routers

import (
	"github.com/gin-gonic/gin"
	"go_server/api"
)

//func (router RouterGroup) SettingRouter() {
//	settingApi := api.ApiGroupApp.SettingsApi
//	router.GET("/", settingApi.SettingsInfoView)
//}

func (ApiRouter) SettingRouter(router *gin.RouterGroup) {
	settingApi := api.ApiGroupApp.SettingsApi
	{
		router.GET("/", settingApi.SettingsInfoView)

	}
}
