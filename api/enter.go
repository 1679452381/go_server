/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package api

import "go_server/api/settings_api"

type AppiGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(AppiGroup)
