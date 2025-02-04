package model

// User 用户表 /*
type User struct {
	Model                   // 基础模型
	Username        string  `gorm:"not null;unique;size: 32;index;type:varchar(32)"` // 用户名
	Password        *string `gorm:"not null;size: 32;type:varchar(127)"`             // 密码
	Avatar          *string // 用户头像
	BackgroundImage *string // 背景图片
	Signature       *string // 个人简介

	updated bool
}
