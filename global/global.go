/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package global

import (
	"go_server/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
