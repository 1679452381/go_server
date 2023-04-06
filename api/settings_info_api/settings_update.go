/**
 * @author X
 * @date 2023/4/6
 * @description
 **/
package settings_info_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_server/config"
	"go_server/core"
	"go_server/global"
	"go_server/models/res"
)

func (SiteInfoApi) SettingsUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsErr, c)
		global.Log.Error(err.Error())
		return
	}
	switch cr.Name {
	case "site":
		var info config.SiteInfo
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParamsErr, c)
			global.Log.Error(err.Error())
			return
		}
		global.Config.SiteInfo = info
		res.OkWithData(global.Config.SiteInfo, c)
	case "qi_niu":
		var qiNiu config.QiNiu
		err := c.ShouldBindJSON(&qiNiu)
		if err != nil {
			res.FailWithCode(res.ParamsErr, c)
			global.Log.Error(err.Error())
			return
		}
		global.Config.QiNiu = qiNiu
		res.OkWithData(global.Config.QiNiu, c)
	default:
		res.FailWithMsg("没有对应配置信息", c)
	}
	//写配置文件
	err = core.SetConf()
	if err != nil {
		fmt.Println(err.Error())
		global.Log.Error(err.Error())
	}
}
