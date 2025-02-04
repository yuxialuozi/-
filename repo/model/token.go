package model

type UserToken struct {
	Token    string `gorm:"not null;primaryKey"`
	Username string `gorm:"not null;unique;type:varchar(32)"` // 用户名，指定 VARCHAR 类型，长度为 32
	UserID   uint32 `gorm:"not null;index"`
	Role     string `gorm:"not null;default:0"` // 用户角色
}
