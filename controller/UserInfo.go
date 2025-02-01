package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/model"
	"simpledouyin/service"
)

func UserInfo(c *gin.Context) {
	var author = model.Author{}

	//获得数据
	userid := c.Query("user_id")
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {

		service.Db.Where("id = ?", userid).Find(&author)

		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "登录成功",
			"user":        author,
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "没有这个用户",
			"user":        author,
		})

	}

}
