/**
 * @author X
 * @date 2023/4/3
 * @description
 **/
package cmd

import (
	"fmt"
	"go_server/core/zlog"
	"go_server/global"
	"go_server/models"
)

// 数据库迁移
func Makemigrations() {
	fmt.Println("开始迁移数据库")

	var err error
	//设置联表
	err = global.DB.SetupJoinTable(&models.AuthModel{}, "CollectsModels", &models.Auth2Collects{})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = global.DB.SetupJoinTable(&models.MenuModel{}, "MenuImages", &models.MenuImageModel{})
	if err != nil {
		fmt.Println(err.Error())
	}
	//获取当前数据库
	//global.DB.Migrator().CurrentDatabase()
	// 生成表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.AuthModel{},
			&models.ArticleModel{},
			&models.ImageModel{},
			&models.TagModel{},
			&models.CateModel{},
			&models.MenuModel{},
			&models.CommentModel{},
			&models.FadeBackMdoel{},
			&models.MassageModel{},
			&models.ModeModel{},
			&models.LoginDataModel{},
		)
	if err != nil {
		zlog.Error("[error ] 创建数据表失败")
		return
	}
	//zlog.Info("[ Info ] 创建数据表完成")
	fmt.Println("创建数据表完成")
}
