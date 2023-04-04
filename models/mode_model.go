/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

type ModeModel struct {
	Model
	Content   string `gorm:"size:128" json:"content"'`
	AuthID    uint
	AuthModel AuthModel `gorm:"foreignKey:AuthID" json:"-"`
}
