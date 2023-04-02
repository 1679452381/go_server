/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

// 自定义菜单和背景图联表
type MenuImageModel struct {
	MenuID     uint       `json:"menu_id"`
	MenuModel  MenuModel  `gorm:"foreignKey:MenuID"`
	ImageID    uint       `json:"image_id"`
	ImageModel ImageModel `gorm:"foreignKey:ImageID"`
	Sort       uint       `json:"sort"`
}
