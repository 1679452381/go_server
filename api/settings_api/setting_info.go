/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package settings_api

import (
	"github.com/gin-gonic/gin"
	"go_server/models/res"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "settingsapi"})
	res.Ok(map[string]interface{}{
		"name": "zxc",
	}, "成功", c)
}
