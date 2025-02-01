package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"simpledouyin/model"
	"simpledouyin/service"
	"strconv"
	"time"
)

/*

不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个

*/

// 视频和time的对应关系
var videoTime = map[int]uint{}

func Feed(c *gin.Context) {
	//实例化对象
	var videolist model.VideoList

	fmt.Println("这个路径运行了一次")

	//获取数值
	latestTime := c.Query("latest_time")
	token := c.Query("token")

	userID := usersLoginInfo[token]

	var author model.Author
	service.Db.Where("id = ?", userID).Find(&author)

	//获取这个时间戳的对象
	Latesttime, _ := strconv.Atoi(latestTime)

	if latestTime == "" {
		Latesttime = int(time.Now().Unix())
	}

	//不是第一次查找
	if Latesttime > 10000000000 {
		Latesttime /= 100
	}

	query := service.Db.Where("create_time < ?", Latesttime).Order("id desc").Limit(2).Preload("Author")
	query.Find(&(videolist.VideoList))

	var minCreateTime int
	if len(videolist.VideoList) > 0 {
		minCreateTime = videolist.VideoList[0].CreateTime
		for _, video := range videolist.VideoList {
			if video.CreateTime < minCreateTime {
				minCreateTime = video.CreateTime
			}
		}
	} else {
		minCreateTime = int(time.Now().Unix())
	}

	//下次会把这个传过来
	videolist.NextTime = minCreateTime

	//遍历判断点赞，如果找到记录，说明点赞过了
	for i := range videolist.VideoList {

		var userLove model.UserLove

		userLove = model.UserLove{
			UserId:  author.ID,
			VideoId: videolist.VideoList[i].ID,
		}

		//如果找到了这个记录，设置为ture
		if service.Db.Where("user_id = ? AND video_id = ?", userLove.UserId, userLove.VideoId).First(&userLove).Error == nil {
			videolist.VideoList[i].IsFavorite = true
		} else {
			videolist.VideoList[i].IsFavorite = false
		}

	}

	c.JSON(200, gin.H{
		"status_code": 0,
		"status_msg":  "获取了视频列表",
		"next_time":   videolist.NextTime,
		"video_list":  videolist.VideoList,
	})

}
