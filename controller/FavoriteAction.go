package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/role"
	"simpledouyin/service"
	"strconv"
)

func FavoriteAction(c *gin.Context) {

	var video role.Video
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	service.Db.Where("id = ?", videoId).Find(&video)

	userID := usersLoginInfo[token]

	var user role.Author
	service.Db.Where("id = ?", userID).Find(&user)

	if _, exist := usersLoginInfo[token]; exist {

		//点赞
		if actionType == "1" {

			//增加这个视频的喜欢数量
			video.FavoriteCount++

			// 增加用户的喜欢作品数量
			user.FavoriteCount++

			var author = role.Author{}
			var authorId uint
			service.Db.Model(&role.Video{}).Where("id = ?", videoId).Pluck("author_id", &authorId)
			service.Db.Model(&role.Author{}).Where("id = ?", authorId).Find(&author)

			//增加作者的点赞数
			intValue, _ := strconv.Atoi(author.TotalFavorited)
			intValue++
			stringValue := strconv.Itoa(intValue)

			// 更新数据库中的喜欢的作品
			service.Db.Model(&role.Author{}).Where("name = ?", user.Name).Update("favorite_count", user.FavoriteCount)
			service.Db.Model(&role.Author{}).Where("name = ?", author.Name).Update("total_favorited", stringValue)

			service.Db.Model(&role.Video{}).Where("id = ?", videoId).Update("favorite_count", video.FavoriteCount)

			var userLove role.UserLove

			userLove = role.UserLove{
				UserId:  user.ID,
				VideoId: video.ID,
			}

			service.Db.Model(&role.UserLove{}).Create(&userLove)

		} else if actionType == "2" {

			//减少这个视频的喜欢数量
			video.FavoriteCount--

			// 减少用户的喜欢作品数量
			user.FavoriteCount--

			var author = role.Author{}
			var authorId uint
			service.Db.Model(&role.Video{}).Where("id = ?", videoId).Pluck("author_id", &authorId)
			service.Db.Model(&role.Author{}).Where("id = ?", authorId).Find(&author)

			//减少作者的点赞数
			intValue, _ := strconv.Atoi(author.TotalFavorited)
			intValue--
			stringValue := strconv.Itoa(intValue)

			// 更新数据库中的作品数量
			service.Db.Model(&role.Author{}).Where("name = ?", user.Name).Update("favorite_count", user.FavoriteCount)
			service.Db.Model(&role.Author{}).Where("name = ?", author.Name).Update("total_favorited", stringValue)

			service.Db.Model(&role.Video{}).Where("id = ?", videoId).Update("favorite_count", video.FavoriteCount)

			var userLove role.UserLove

			userLove = role.UserLove{
				UserId:  user.ID,
				VideoId: video.ID,
			}

			service.Db.Where("user_id = ? AND video_id = ?", userLove.UserId, userLove.VideoId).Delete(&role.UserLove{})

		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "成功",
		})

	} else {
		// 未找到记录的处理逻辑
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "没有这个token",
		})
	}

}
