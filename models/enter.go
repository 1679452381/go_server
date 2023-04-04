/**
 * @author X
 * @date 2023/3/31
 * @description
 **/
package models

import "time"

type Model struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"-"`
}
