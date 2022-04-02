package models

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type TUser struct {
	ID        uint        `json:"id,omitempty" gorm:"primary_key"`
	Account   string      `json:"account" gorm:"type:varchar;not null"`
	Password  string      `json:"password" gorm:"type:varchar;not null"`
	RoleID    null.Int    `json:"role_id,omitempty" gorm:"type:tinyint"`
	Nickname  null.String `json:"nickname,omitempty" gorm:"type:varchar"`
	Avatar    null.String `json:"avatar,omitempty" gorm:"type:varchar"`
	Country   null.String `json:"country,omitempty" gorm:"type:varchar"`
	Province  null.String `json:"province,omitempty" gorm:"type:varchar"`
	City      null.String `json:"city,omitempty" gorm:"type:varchar"`
	Language  null.String `json:"language,omitempty" gorm:"type:varchar"`
	Phone     null.String `json:"phone,omitempty" gorm:"type:varchar"`
	EMail     null.String `json:"e_mail,omitempty" gorm:"type:varchar"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	DeletedAt time.Time   `json:"deleted_at,omitempty"`
}

type User struct {
	ID        uint        `json:"id,omitempty" gorm:"primary_key"`
	RoleID    null.Int    `json:"role_id,omitempty" gorm:"type:tinyint"`
	Nickname  null.String `json:"nickname,omitempty" gorm:"type:varchar"`
	Avatar    null.String `json:"avatar,omitempty" gorm:"type:varchar"`
	Country   null.String `json:"country,omitempty" gorm:"type:varchar"`
	Province  null.String `json:"province,omitempty" gorm:"type:varchar"`
	City      null.String `json:"city,omitempty" gorm:"type:varchar"`
	Language  null.String `json:"language,omitempty" gorm:"type:varchar"`
	Phone     null.String `json:"phone,omitempty" gorm:"type:varchar"`
	EMail     null.String `json:"e_mail,omitempty" gorm:"type:varchar"`
}
