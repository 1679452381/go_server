/**
 * @author X
 * @date 2023/4/1
 * @description
 **/
package models

type ImageModel struct {
	Model
	Path string `json:"path"`                //图片路径
	Hash string `json:"hash"`                //图片hash值
	Name string `gorm:"size:38" json:"name"` //图片名称
}
