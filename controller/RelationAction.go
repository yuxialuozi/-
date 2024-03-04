package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"simpledouyin/role"
	"simpledouyin/service"
)

func RelationAction(c *gin.Context) {
	var toUser role.Author

	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")

	//获取要关注的人
	service.Db.Model(&role.Author{}).Where("id = ?", toUserId).Find(&toUser)

	if _, exist := usersLoginInfo[token]; exist {
		//1-关注，2-取消关注

		userID := usersLoginInfo[token]

		var user role.Author
		service.Db.Where("id = ?", userID).Find(&user)

		if user.Name == toUser.Name {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  "不可以关注自己",
			})
			return
		}

		var relation role.Relation
		err := service.Db.Where("user_id = ? AND to_user_id = ?", user.ID, toUser.ID).First(&relation).Error

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 1,
				"status_msg":  "你已经关注过对方了",
			})
			return
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			if actionType == "1" {

				// 获取用户信息
				userID := usersLoginInfo[token]

				var user role.Author

				service.Db.Where("id = ?", userID).Find(&user)

				//增加这个user的关注数量
				user.FollowCount++

				// 增加关注人的被关注数量
				toUser.FollowerCount++

				//更新和建立表关系
				service.Db.Model(&role.Author{}).Where("name = ?", user.Name).Update("follow_count", user.FollowCount)
				service.Db.Model(&role.Author{}).Where("name = ?", toUser.Name).Update("follower_count", toUser.FollowerCount)

				var relation role.Relation
				relation = role.Relation{
					UserId:   user.ID,
					ToUserId: toUser.ID,
				}

				service.Db.Model(&role.Relation{}).Create(&relation)

			} else if actionType == "2" {

				//获得user
				userID := usersLoginInfo[token]

				var user role.Author
				service.Db.Where("id = ?", userID).Find(&user)

				//减少这个user的关注数量
				user.FollowCount--

				// 减少关注人的被关注数量
				toUser.FollowerCount--

				var relation role.Relation
				relation.UserId = user.ID
				relation.ToUserId = toUser.ID

				//更新和建立表关系
				service.Db.Model(&role.Video{}).Where("name = ?", user.Name).Update("follow_count", user.FollowCount)
				service.Db.Model(&role.Author{}).Where("name = ?", toUser.Name).Update("follower_count", toUser.FollowerCount)

				relation = role.Relation{
					UserId:   user.ID,
					ToUserId: toUser.ID,
				}

				service.Db.Where("user_id = ? AND to_user_id = ?", relation.UserId, relation.ToUserId).Delete(&role.Relation{})

			}

			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "成功",
			})

		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "没有登录，拒绝关注",
		})
	}

}
