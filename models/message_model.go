/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

type MassageModel struct {
	Model
	SendAuthID       uint      `gorm:"primaryKey" json:"send-auth_id"`
	SendAuthModel    AuthModel `gorm:"foreignKey:SendAuthID" json:"-"`
	SendAuthNickName string    `gorm:"size:42" json:"send_auth_nick_name"`
	SendAuthAvatar   string    ` json:"send_auth_avatar"`

	RevAuthID       uint      `gorm:"primaryKey" json:"rev_auth_id"`
	RevAuthModel    AuthModel `gorm:"foreignKey:RevAuthID" json:"-"`
	RevAuthNickName string    `gorm:"size:42" json:"rev_auth_nick_name"`
	RevAuthAvatar   string    ` json:"rev_auth_avatar"`
	IsRead          bool      `gorm:"default:false" json:"is_read"`
	Content         string    `json:"content"`
}
