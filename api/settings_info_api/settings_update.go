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
	var conf config.SiteInfo
	err := c.ShouldBindJSON(&conf)
	if err != nil {
		res.FailWithCode(res.ParamsErr, c)
		global.Log.Error(err.Error())
		return
	}

	global.Config.SiteInfo = conf
	//写配置文件
	err = core.SetConf()
	if err != nil {
		fmt.Println(err.Error())
		global.Log.Error(err.Error())
	}
	res.Ok(conf, "配置成功", c)
}
