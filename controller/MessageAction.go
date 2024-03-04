package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/role"
	"simpledouyin/service"
	"strconv"
	"time"
)

func MessageAction(c *gin.Context) {
	// 获取数据
	token := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	toUserIdInt, err := strconv.Atoi(toUserIdStr)
	actionType := c.Query("action_type")
	content := c.Query("content")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": 1,
			"status_msg":  "无效的 to_user_id",
		})
		return
	}

	if _, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			// 获得user对象

			userID := usersLoginInfo[token]

			var user role.Author
			service.Db.Where("id = ?", userID).Find(&user)

			message := role.Message{
				ToUserID:   uint(toUserIdInt),
				FromUserID: user.ID,
				Content:    content,
				CreateTime: int(time.Now().Unix()),
				IsViewed:   true,
			}

			if err := service.Db.Model(&role.Message{}).Create(&message).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 2,
					"status_msg":  "发送失败",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "发送成功",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "没有发送",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "你还没有登录",
		})
	}
}
