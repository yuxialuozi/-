package role

type UserLove struct {
	ID      uint `json:"id" gorm:"primary_key" `
	UserId  uint
	VideoId uint
}
