package model

import "time"

type Message struct {
	Id          int64     `json:"id" gorm:"id"`                     // 主键ID
	Uid         string    `json:"uid" gorm:"uid"`                   //uid
	UserId      int64     `json:"user_id" gorm:"user_id"`           // 用户id
	ToUserId    int64     `json:"to_user_id" gorm:"to_user_id"`     // 客户id
	RoomId      int64     `json:"room_id" gorm:"room_id"`           // 0为私聊
	Content     string    `json:"content" gorm:"content"`           // 聊天内容
	ImageUrl    string    `json:"image_url" gorm:"image_url"`       // 图片url
	CreatedTime time.Time `json:"created_time" gorm:"created_time"` // 创建时间
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"` // 更新时间
}

// TableName 表名称
func (*Message) TableName() string {
	return "message"
}
