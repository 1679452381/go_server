/**
 * @author X
 * @date 2023/4/6
 * @description
 **/
package routers

import (
	"github.com/gin-gonic/gin"
	"go_server/api"
)

func (ApiRouter) ImagesRouter(router *gin.RouterGroup) {
	imageApi := api.ApiGroupApp.ImageApi
	{
		router.POST("/images", imageApi.ImageUploadView)
	}
}
