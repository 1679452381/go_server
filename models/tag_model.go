/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

type TagModel struct {
	Model
	Title         string      `gorm:"size:16" json:"title"`           //标签名称
	ArticleModels []AuthModel `gorm:"many2many:article_tag" json:"-"` //对应文章
}
