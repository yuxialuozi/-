package Router

import (
	"github.com/gin-gonic/gin"
	"simpledouyin/controller"
)

func InitRouter(r *gin.Engine) {

	// 设置静态文件路径
	r.Static("/public", "./public")

	apiRouter := r.Group("/douyin") //分组

	// basic apis

	apiRouter.GET("/feed/", controller.Feed)                //完成
	apiRouter.GET("/user/", controller.UserInfo)            //完成
	apiRouter.POST("/user/register/", controller.Register)  //完成
	apiRouter.POST("/user/login/", controller.Login)        //完成
	apiRouter.POST("/publish/action/", controller.Publish)  //完成
	apiRouter.GET("/publish/list/", controller.PublishList) //完成

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction) //完成
	apiRouter.GET("/favorite/list/", controller.FavoriteList)      //完成
	apiRouter.POST("/comment/action/", controller.CommentAction)   //完成
	apiRouter.GET("/comment/list/", controller.CommentList)        //完成

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)     //完成
	apiRouter.GET("/relation/follow/list/", controller.FollowList)     //完成
	apiRouter.GET("/relation/follower/list/", controller.FollowerList) //完成
	apiRouter.GET("/relation/friend/list/", controller.FriendList)     //完成
	apiRouter.GET("/message/chat/", controller.MessageChat)            //完成
	apiRouter.POST("/message/action/", controller.MessageAction)       //完成
}
