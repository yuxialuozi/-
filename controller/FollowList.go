package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/model"
	"simpledouyin/service"
)

func FollowList(c *gin.Context) {
	var followList []model.Author

	userId := c.Query("user_id")
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		var followid []uint

		service.Db.Model(&model.Relation{}).Where("user_id = ?", userId).Pluck("to_user_id", &followid)

		fmt.Println(followid)
		//获得了他关注的所有人

		for _, id := range followid {
			var follow model.Author

			service.Db.Model(&model.Author{}).Where("id = ?", id).First(&follow)

			followList = append(followList, follow)
		}

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "找到了",
			"user_list":   followList,
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "你还没有登录",
		})

	}

}
