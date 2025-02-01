package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simpledouyin/model"
	"simpledouyin/service"
)

// 把token和用户的id进行绑定
var usersLoginInfo = map[string]uint{}

func Register(c *gin.Context) {
	//注册用户，第一条代码进行实例化
	var author = model.Author{}

	//获取信息
	username := c.Query("username")
	password := c.Query("password")

	//按照规则生成token
	token := username + password

	if _, exist := usersLoginInfo[token]; exist {

		//返回响应信息
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "注册失败，用户名重复",
		})
	} else {

		//赋值并且创建记录
		author.Name = username
		author.PassWord = password

		service.Db.Create(&author)

		//加入记录
		usersLoginInfo[token] = author.ID

		//返回响应信息
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "注册成功",
			"user_id":     author.ID,
			"token":       token,
		})

	}

}
