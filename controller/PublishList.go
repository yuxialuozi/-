package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/role"
	"simpledouyin/service"
)

func PublishList(c *gin.Context) {
	//实例化对象
	var author = role.Author{}
	var videolist role.VideoList

	//获取参数
	token := c.Query("token")
	userId := c.Query("user_id")

	if _, exist := usersLoginInfo[token]; exist {

		service.Db.Where("id = ?", userId).Find(&author)

		service.Db.Where("author_id = ?", author.ID).Find(&(videolist.VideoList))

		//遍历绑定作者
		for i := range videolist.VideoList {
			videolist.VideoList[i].Author = author
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
