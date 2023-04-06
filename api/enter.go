/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package api

import (
	"go_server/api/settings_info_api"
)

type AppiGroup struct {
	SiteInfoApi settings_info_api.SiteInfoApi
}

var ApiGroupApp = new(AppiGroup)
