/**
 * @author X
 * @date 2023/3/30
 * @description  连接数据库
 **/
package core

import (
	"fmt"
	"go_server/core/zlog"
	"go_server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGrom() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		zlog.Warn("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		//	开发环境显示所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)

	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: mysqlLogger})
	if err != nil {
		zlog.Error(fmt.Sprintf("[%s]连接失败", db))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxIdleTime(time.Hour * 4) //连接最大复用时间
	return db
}
