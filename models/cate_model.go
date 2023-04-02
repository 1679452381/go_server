/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

type CateModel struct {
	Model
	Title         string      `gorm:"size:16" json:"title"`           //分类名称
	ArticleModels []AuthModel `gorm:"many2many:article_tag" json:"-"` //对应文章
}
