package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/role"
	"simpledouyin/service"
)

func MessageChat(c *gin.Context) {
	var messageList []role.Message

	token := c.Query("token")
	toUserID := c.Query("to_user_id")

	userID := usersLoginInfo[token]

	var user role.Author
	service.Db.Where("id = ?", userID).Find(&user)

	// 检查用户是否已登录
	if _, exist := usersLoginInfo[token]; exist {
		// 查询未被查看的消息记录，并按照创建时间降序排序
		err := service.Db.Model(&role.Message{}).
			Where("((to_user_id = ? AND from_user_id = ?) OR (to_user_id = ? AND from_user_id = ?)) AND is_viewed = ?",
				toUserID, user.ID, user.ID, toUserID, false).
			Find(&messageList).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "查询聊天记录失败",
			})
			return
		}

		// 标记已查看的消息
		for _, message := range messageList {
			message.IsViewed = true

			service.Db.Save(&message)
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code":  0,
			"status_msg":   "找到了聊天记录",
			"message_list": messageList,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status_code": 1,
			"status_msg":  "你还没有登录",
		})
	}
}
