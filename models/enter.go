/**
 * @author X
 * @date 2023/3/31
 * @description
 **/
package models

import "time"

type Model struct {
	ID       uint      `gorm:"primaryKey" json:"id" :"id"`
	CreateAt time.Time `json:"create_at" :"create-at"`
	UpdateAt time.Time `json:"-" :"update-at"`
}
