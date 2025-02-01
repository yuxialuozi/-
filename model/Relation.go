package model

type Relation struct {
	ID       uint `json:"id" gorm:"primary_key" ` //主键id
	UserId   uint
	ToUserId uint `gorm:""`
}
