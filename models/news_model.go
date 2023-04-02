/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

type FadeBackMdoel struct {
	Model
	Email        string `gorm:"size:64" json:"email"`
	Content      string `gorm:"size:128" json:"content"`
	ApplyContent string `gorm:"size:128" json:"apply_content"`
	IsApply      bool   `gorm:"size:2" json:"is_apply"`
}
