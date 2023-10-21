package models

import "time"

// 数据库对应的实体
type User struct {
	Uid      int       `json:"uid" orm:"uid"`            // 用户ID
	UserName string    `json:"userName" orm:"user_name"` // 用户名
	Passwd   string    `json:"passwd" orm:"passwd"`      // 用户密码（请注意，存储密码时应该使用哈希而不是明文）
	Avatar   string    `json:"avatar" orm:"avatar"`      // 用户头像
	CreateAt time.Time `json:"createAt" orm:"create_at"` // 用户创建时间
	UpdateAt time.Time `json:"updateAt" orm:"update_at"` // 用户更新时间
}

// 需要的部分数据的实体
type UserInfo struct {
	Uid      int    `json:"uid"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
