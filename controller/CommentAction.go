package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/model"
	"simpledouyin/service"
	"strconv"
	"time"
)

func CommentAction(c *gin.Context) {

	var video model.Video
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")

	commentId := c.Query("comment_id")

	//获取video
	service.Db.Where("id = ?", videoId).Find(&video)

	// 获取评论用户信息
	userID := usersLoginInfo[token]

	var user model.Author
	service.Db.Where("id = ?", userID).Find(&user)

	if _, exist := usersLoginInfo[token]; exist {
		//1-发布评论，2-删除评论
		if actionType == "1" {

			var comment model.Comment
			comment.User = user
			comment.UserID = user.ID
			videouintid, _ := strconv.ParseUint(videoId, 10, 64)
			comment.VideoID = uint(videouintid)

			var video model.Video
			service.Db.Model(&model.Video{}).Where("id = ?", videouintid).Find(&video)

			video.CommentCount++

			service.Db.Model(&model.Video{}).Where("id = ?", videouintid).Update("comment_count", video.CommentCount)

			comment.Video = video

			comment.Content = commentText

			now := time.Now()
			dateString := now.Format("01-02")

			comment.CreateDate = dateString

			service.Db.Model(&model.Comment{}).Create(&comment)

			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "增加了用户评论",
				"comment":     comment,
			})

		} else if actionType == "2" {

			var comment model.Comment
			service.Db.Model(&model.Comment{}).Where("id = ?", commentId).Delete(&comment)

			videouintid, _ := strconv.ParseUint(videoId, 10, 64)
			comment.VideoID = uint(videouintid)

			video.CommentCount--
			service.Db.Model(&model.Video{}).Where("id = ?", videouintid).Update("comment_count", video.CommentCount)

			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "删除了用户评论",
				"comment":     comment,
			})

		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "没有登录，拒绝评论",
		})
	}

}
