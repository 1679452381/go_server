/**
 * @author X
 * @date 2023/4/1
 * @description 用户收藏表
 **/
package models

import "time"

// 记录用户什么时候收藏了什么文章
type Auth2Collects struct {
	AuthID       uint         `gorm:"primaryKey" json:"auth_id"`
	AuthModel    AuthModel    `gorm:"foreignKey:AuthID" json:"auth_model"`
	ArticleID    uint         `gorm:"primaryKey" json:"article_id"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"article_model"`
	CreateAt     time.Time    `json:"create_at"`
}
