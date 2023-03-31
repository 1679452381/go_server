/**
 * @author X
 * @date 2023/3/31
 * @description
 **/
package models

import "go_server/models/ctype"

type AuthModel struct {
	Model
	NickName       string           `gorm:"size:42" json:"nick_name"`
	UserName       string           `gorm:"size:36" json:"user_name"`
	Password       string           `gorm:"size:128" json:"password"`
	Avatar         string           `gorm:"size:256" json:"avatar"`
	Email          string           `gorm:"size:128"json:"email"`
	Tel            string           `gorm:"size:18" json:"tel"`
	Addr           string           `gorm:"size:64" json:"addr"`
	Token          string           `gorm:"size:64" json:"token"`               //其他平台唯一id
	IP             string           `gorm:"size:20" json:"IP"`                  //ip
	Role           ctype.Role       `gorm:"size:4;default:1" json:"role"`       // 权限1 管理员 2 普通用户 3 游客
	SignStatus     ctype.SignStatus `gorm:"type=smallint(6)" json:"signStatus"` //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignKey:AuthID" json:"articleModels"`
	CollectsModels []ArticleModel   `gorm:"many2many:auth2_collects;JoinForeignKey:AuthID;JoinReference:ArticleID" json:"collectsModels"`
}
