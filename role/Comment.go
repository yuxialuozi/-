package role

type Comment struct {
	ID         uint   `json:"id" gorm:"primary_key" ` //评论id
	User       Author `json:"user" gorm:"foreignKey:UserID"`
	UserID     uint   `json:"-" gorm:"index"`
	Video      Video  `json:"-" gorm:"foreignKey:VideoID"`
	VideoID    uint   `json:"-" gorm:"index"`
	Content    string `json:"content"`     //评论内容
	CreateDate string `json:"create_date"` //评论发布日期
}
