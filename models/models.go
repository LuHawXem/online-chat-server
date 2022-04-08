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
	Gender	  null.Int	  `json:"gender,omitempty" gorm:"tinyint" remark:"性别,0男,1女"`
	Nickname  null.String `json:"nickname,omitempty" gorm:"varchar"`
	Avatar    null.String `json:"avatar,omitempty" gorm:"varchar"`
	Country   null.String `json:"country,omitempty" gorm:"varchar"`
	Province  null.String `json:"province,omitempty" gorm:"varchar"`
	City      null.String `json:"city,omitempty" gorm:"varchar"`
	Language  null.String `json:"language,omitempty" gorm:"varchar"`
	Phone     null.String `json:"phone,omitempty" gorm:"varchar"`
	EMail     null.String `json:"e_mail,omitempty" gorm:"varchar"`
	Status    uint8       `json:"status,omitempty" gorm:"tinyint;default:0" remark:"数据的状态,0表示存在,1表示已删除,可扩展,不使用DeletedAt字段是为了方便在数据库中删除数据"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
}

type User struct {
	ID       uint32      `json:"id,omitempty" gorm:"primary_key"`
	Account  uint64      `json:"account" gorm:"bigint"`
	RoleID   uint8       `json:"role_id,omitempty" gorm:"tinyint;default:0" remark:"0表示普通用户"`
	Gender	  null.Int	  `json:"gender,omitempty" gorm:"tinyint" remark:"性别,0男,1女"`
	Nickname null.String `json:"nickname,omitempty" gorm:"varchar"`
	Avatar   null.String `json:"avatar,omitempty" gorm:"varchar"`
	Country  null.String `json:"country,omitempty" gorm:"varchar"`
	Province null.String `json:"province,omitempty" gorm:"varchar"`
	City     null.String `json:"city,omitempty" gorm:"varchar"`
	Status   uint8       `json:"status,omitempty" gorm:"tinyint;default:0" remark:"数据的状态,0表示存在,1表示已删除,可扩展"`
}

type TRelation struct {
	ID        uint32      `json:"id,omitempty" gorm:"primary_key"`
	UserID    uint32      `json:"user_id,omitempty" gorm:"int" remark:"关系中对方的UserID"`
	Remark    null.String `json:"remark,omitempty" gorm:"varchar" remark:"关系中对对方的备注"`
	CreatedBy uint32      `json:"created_by,omitempty" gorm:"int" remark:"关系中己方的UserID"`
	State     uint8       `json:"state" gorm:"tinyint;default:1" remark:"关系的状态,0表示双向已建立,1表示单向(己方->对方)已建立"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	Status 	  uint8		  `json:"status,omitempty" gorm:"tinyint;default:0" remark:"数据的状态,0表示存在,1表示已删除,可扩展"`
}

type TMessage struct {
	ID        uint32      `json:"id,omitempty" gorm:"primary_key"`
	Receiver  uint32	  `json:"receiver,omitempty" gorm:"int" remark:"消息接收方id"`
	ReplyID	  null.Int    `json:"reply_id,omitempty" gorm:"int" remark:"回复消息的id,可能没有"`
	Operate	  null.Int	  `json:"operate,omitempty" gorm:"tinyint" remark:"操作,如回复消息(存在ReplyID且不存在Operate即为回复),通过请求(1)与否(0)"`
	Content   null.String `json:"content,omitempty" gorm:"varchar" remark:"正文,可能没有"`
	Type      uint8 	  `json:"type" gorm:"tinyint" remark:"消息类型,0普通消息(对单),1加好友(对单),2群消息,3加群"`
	State	  uint8		  `json:"state,omitempty" gorm:"tinyint" remark:"消息状态,0已发送,1已接收,2已处理"`
	CreatedBy uint32	  `json:"created_by,omitempty" gorm:"int"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	Status 	  uint8		  `json:"status,omitempty" gorm:"tinyint;default:0" remark:"数据的状态,0表示存在,1表示已删除,可扩展"`
}

type Message struct {
	Receiver  uint32	  `json:"receiver,omitempty" remark:"消息接收方id"`
	ReplyID	  null.Int    `json:"reply_id,omitempty" remark:"回复消息的id,可能没有"`
	Operate	  null.Int	  `json:"operate,omitempty" remark:"操作,如回复消息(存在ReplyID且不存在Operate即为回复),通过请求(1)与否(0)"`
	Content   null.String `json:"content,omitempty" remark:"正文,可能没有"`
	Type      uint8 	  `json:"type" remark:"消息类型,0普通消息(对单),1加好友(对单),2群消息,3加群,4搜索用户,5获取好友列表,200心跳包"`
}