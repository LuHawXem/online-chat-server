package models

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type TUser struct {
	ID        uint32      `json:"id,omitempty" gorm:"primary_key"`
	Account   uint64      `json:"account" gorm:"bigint"`
	Password  string      `json:"password" gorm:"varchar"`
	RoleID    uint8       `json:"role_id,omitempty" gorm:"tinyint;default:0" remark:"0表示普通用户"`
	Nickname  null.String `json:"nickname,omitempty" gorm:"varchar"`
	Avatar    null.String `json:"avatar,omitempty" gorm:"varchar"`
	Country   null.String `json:"country,omitempty" gorm:"varchar"`
	Province  null.String `json:"province,omitempty" gorm:"varchar"`
	City      null.String `json:"city,omitempty" gorm:"varchar"`
	Language  null.String `json:"language,omitempty" gorm:"varchar"`
	Phone     null.String `json:"phone,omitempty" gorm:"varchar"`
	EMail     null.String `json:"e_mail,omitempty" gorm:"varchar"`
	Status    uint8       `json:"status,omitempty" gorm:"tinyint;default:0" remark:"0表示存在, 1表示已删除"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
}

type User struct {
	ID       uint32      `json:"id,omitempty" gorm:"primary_key"`
	RoleID   uint8       `json:"role_id,omitempty" gorm:"tinyint;default:0" remark:"0表示普通用户"`
	Nickname null.String `json:"nickname,omitempty" gorm:"varchar"`
	Avatar   null.String `json:"avatar,omitempty" gorm:"varchar"`
	Country  null.String `json:"country,omitempty" gorm:"varchar"`
	Province null.String `json:"province,omitempty" gorm:"varchar"`
	City     null.String `json:"city,omitempty" gorm:"varchar"`
	Language null.String `json:"language,omitempty" gorm:"varchar"`
	Phone    null.String `json:"phone,omitempty" gorm:"varchar"`
	EMail    null.String `json:"e_mail,omitempty" gorm:"varchar"`
	Status   uint8       `json:"status,omitempty" gorm:"tinyint;default:0" remark:"0表示存在, 1表示已删除"`
}
