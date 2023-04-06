/**
 * @author X
 * @date 2023/4/5
 * @description
 **/
package settings_info_api

import (
	"github.com/gin-gonic/gin"
	"go_server/global"
	"go_server/models/res"
)

func (SiteInfoApi) SiteInfoView(c *gin.Context) {
	//fmt.Println(global.Config.SiteInfo)
	res.OkWithData(global.Config.SiteInfo, c)
}
