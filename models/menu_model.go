/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

import "go_server/models/ctype"

type MenuModel struct {
	Model
	ModelTitle   string       `gorm:"size:32" json:"model_title"`
	ModelTitleEn string       `gorm:"size:32" json:"model_title_en"`
	Slogan       string       `gorm:"size:64" json:"slogan"`                                                                        //slo gan
	Abstract     ctype.Array  `gorm:"type:string" json:"abstract"`                                                                  //简介
	AbstractTime int          `json:"abstract_time"`                                                                                //简介切换时间
	MenuImages   []ImageModel `gorm:"many2many:menu_image_models;JoinForeignKey:MenuID;JoinReferences:ImageID " json:"menu_images"` //菜单的图片列表
	MenuTime     int          `json:"menu_time"`                                                                                    //菜单的切换时间 0为不切换
	Sort         int          `gorm:"size:10" json:"sort"`                                                                          //菜单的顺序
}
