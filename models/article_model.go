/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

import "go_server/models/ctype"

type ArticleModel struct {
	Model
	Tile          string         `gorm:"size:32" json:"tile"`               //标题
	Abstract      string         `json:"abstract"`                          //摘要
	Content       string         `json:"content"`                           // 内容
	LookCount     int            `json:"look_count"`                        // 浏览量
	CommentCount  int            `json:"comment_count"`                     //评论数
	DiggCount     int            `json:"digg_count"`                        // 点赞数
	CollectCount  int            `json:"collect_count"`                     //被收藏数
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"-"`    //文章标签
	Tags          ctype.Array    `json:"tags"`                              //文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"-"`     //文章评论
	AuthModel     AuthModel      `gorm:"many2many:auth2_collects" json:"-"` //文章作者
	AuthID        uint           `json:"auth_id"`                           //作者id
	NickName      string         `gorm:"size:42" json:"nick_name"`          //作者昵称
	CateModel     CateModel      `gorm:"foreignKey:CateID"`
	CateID        uint           `json:"cate_id"`                     //文章分类
	Category      string         `gorm:"size:16"  json:"category"`    //文章分类
	Source        string         `json:"source"`                      //文章来源
	Link          string         `json:"link"`                        //文章链接
	Image         ImageModel     `gorm:"foreignKey:ImageID" json:"-"` //文章封面
	ImageID       uint           `json:"image_id"`                    //文章封面
	ImagePath     string         `json:"image_path"`
}
