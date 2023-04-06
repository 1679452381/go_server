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

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SiteInfoApi) SiteInfoView(c *gin.Context) {
	//fmt.Println(global.Config.SiteInfo)
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsErr, c)
		return
	}
	switch cr.Name {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "qi_niu":
		res.OkWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMsg("没有对应配置信息", c)
	}

}
