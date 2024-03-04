package role

type Message struct {
	ID         uint   `json:"id" gorm:"primary_key" ` //主键id
	ToUserID   uint   `json:"to_user_id"`             //信息接收者
	FromUserID uint   `json:"from_user_id"`           //信息发送者
	Content    string `json:"content"`                //信息内容
	CreateTime int    `json:"create_time"`            //信息发送时间
	IsViewed   bool   `json:"-" gorm:"default:false"` //是否已经找到了这个消息
}
