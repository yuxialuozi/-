package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/role"
	"simpledouyin/service"
)

func Login(c *gin.Context) {
	var author = role.Author{}

	//获取信息
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	err := service.Db.Where("name = ? AND pass_word = ?", username, password).First(&author).Error

	if err == nil {

		usersLoginInfo[token] = author.ID

		//返回响应信息
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "登录成功",
			"user_id":     author.ID,
			"token":       token,
		})

	} else {
		// 未找到记录的处理逻辑
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "登录失败，用户名或密码错误",
		})
	}
}
