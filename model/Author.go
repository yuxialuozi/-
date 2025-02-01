package model

type Author struct {
	ID              uint   `json:"id" gorm:"primary_key" `                                                               //用户id
	Name            string `json:"name" gorm:"unique;not null"`                                                          //用户名称
	PassWord        string `json:"-"`                                                                                    //用户密码
	FollowCount     int    `json:"follow_count"`                                                                         //关注总数
	FollowerCount   int    `json:"follower_count"`                                                                       //粉丝总数
	IsFollow        bool   `json:"is_follow"`                                                                            //是否关注，ture已经关注，false没有关注
	Avatar          string `json:"avatar" gorm:"default:'http://192.168.1.4:8080/public/picture/morentouxiang.png'"`     //用户头像
	BackgroundImage string `json:"background_image" gorm:"default:'http://192.168.1.4:8080/public/picture/beijing.png'"` //用户个人页顶部大图
	Signature       string `json:"signature" gorm:"default:'这个人很懒，还没有写个人简介'"`                                            //个人简介
	TotalFavorited  string `json:"total_favorited" gorm:"default:'0'"`                                                   //获赞总数
	WorkCount       int    `json:"work_count"`                                                                           //作品数
	FavoriteCount   int    `json:"favorite_count"`                                                                       //喜欢数
}
