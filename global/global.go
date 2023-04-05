/**
 * @author X
 * @date 2023/3/30
 * @description
 **/
package global

import (
	"go.uber.org/zap"
	"go_server/config"
	"gorm.io/gorm"
)

var (
	Config  *config.Config
	DB      *gorm.DB
	Zlogger *zap.Logger
)
