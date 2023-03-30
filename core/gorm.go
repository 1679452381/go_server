/**
 * @author X
 * @date 2023/3/30
 * @description  连接数据库
 **/
package core

import (
	"fmt"
	"go_server/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitGrom() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("未配置mysql")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	if global.Config.System.Env == "debug" {
		//	开发环境显示所有sql
	} else {

	}

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf(fmt.Sprintf("[%s]连接失败"), db)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxIdleTime(time.Hour * 4) //连接最大复用时间
	return db
}
