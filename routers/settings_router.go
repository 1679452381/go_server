/**
 * @author X
 * @date 2023/4/5
 * @description
 **/
package routers

import (
	"github.com/gin-gonic/gin"
	"go_server/api"
)

func (ApiRouter) SiteInfoRouter(router *gin.RouterGroup) {
	siteInfoApi := api.ApiGroupApp.SiteInfoApi
	{
		router.GET("/settings/:name", siteInfoApi.SiteInfoView)
		router.PUT("/settings/:name", siteInfoApi.SettingsUpdateView)
	}
}
