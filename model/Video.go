package model

type Video struct {
	ID            uint   `json:"id" gorm:"primary_key"`             //视频唯一标识
	Author        Author `json:"author" gorm:"ForeignKey:AuthorID"` //视频作者信息,外键关联到 User 结构体的 ID 字段
	AuthorID      uint   `json:"-"`
	PlayUrl       string `json:"play_url" `                       //视频播放地址
	CoverUrl      string `json:"cover_url" `                      //视频封面地址
	FavoriteCount int    `json:"favorite_count" gorm:"default:0"` //视频的点赞总数
	CommentCount  int    `json:"comment_count" gorm:"default:0"`  //视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`                     //是否点赞，ture已经点赞，false没有点赞
	Title         string `json:"title"`                           //视频标题
	CreateTime    int    `json:"-"`                               //视频创建的时间戳
}
