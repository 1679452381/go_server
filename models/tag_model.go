/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

type TagModel struct {
	Model
	Title         string         `gorm:"size:16" json:"title"`                                                                  //标签名称
	ArticleModels []ArticleModel `gorm:"many2many:article_tag_models;joinForeignKey:tag_id;joinReferences:article_id" json:"-"` //对应文章
}
