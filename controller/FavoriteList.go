package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/model"
	"simpledouyin/service"
)

func FavoriteList(c *gin.Context) {

	var videolist model.VideoList

	//获取参数
	token := c.Query("token")
	userId := c.Query("user_id")

	if _, exist := usersLoginInfo[token]; exist {

		var videoid []uint

		service.Db.Model(&model.UserLove{}).Where("user_id = ?", userId).Pluck("video_id", &videoid)

		for _, id := range videoid {
			var video model.Video
			var authorId uint
			var author = model.Author{}

			service.Db.First(&video, id)

			service.Db.Model(&model.Video{}).Where("id = ?", id).Pluck("author_id", &authorId)

			service.Db.Model(&model.Author{}).Where("id = ?", authorId).Find(&author)

			video.Author = author
			videolist.VideoList = append(videolist.VideoList, video)
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "找到了",
			"video_list":  videolist.VideoList,
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "没有这个用户",
		})

	}

}
