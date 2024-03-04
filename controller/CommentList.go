package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/role"
	"simpledouyin/service"
	"strconv"
)

func CommentList(c *gin.Context) {
	var commentList []role.Comment
	token := c.Query("token")
	videoId := c.Query("video_id")

	fmt.Println(token)

	videouintid, _ := strconv.ParseUint(videoId, 10, 64)

	var userId []uint

	service.Db.Model(&role.Comment{}).Where("video_id = ?", videouintid).Order("id desc").Pluck("user_id", &userId)
	fmt.Println(userId)

	service.Db.Model(&role.Comment{}).Where("video_id = ?", videouintid).Order("id desc").Find(&commentList)

	for i := range userId {
		var user role.Author

		service.Db.Model(&role.Author{}).Where("id = ?", userId[i]).Find(&user)

		commentList[i].User = user
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   "找到了",
		"comment_list": commentList,
	})

}
