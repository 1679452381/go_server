/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

type CommentModel struct {
	Model
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`  //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"` //父评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                               //父评论id

	Content      string       `gorm:"size:256" json:"content"`               //评论内容
	DiggCount    int          `gorm:"size:8,default:0" json:"digg_count"`    //点赞数
	CommentCount int          `gorm:"size:8,default:0" json:"comment_count"` //评论数
	Article      ArticleModel `gorm:"foreignKey:ArticleID" json:"article"`   //关联 文章
	ArticleID    uint         ` json:"article_id"`                           //文章id
	Auth         AuthModel    `gorm:"foreignKey:AuthID" json:"auth"`         //用户
	AuthID       uint         ` json:"auth_id"`                              //用户id
	NickName     string       `gorm:"size:42" json:"nick_name"`              // 用户昵称
	Avatar       string       `json:"avatar"`                                //用户头像
}
