package main

import (
	"github.com/gin-gonic/gin"
	"simpledouyin/Router"
	"simpledouyin/config"
	"simpledouyin/logging"
	"simpledouyin/migrations" // 引入 migrations 包
	"simpledouyin/service"
)

func main() {
	// 初始化日志配置
	logging.Init()

	// 初始化数据库连接
	DB, err := config.InitDB()
	if err != nil {
		logging.Logger.Fatalf("failed to initialize database: %v", err)
	}

	// 执行数据库自动迁移
	migrations.AutoMigrate(DB)

	// 将数据库连接赋值给服务
	service.Db = DB
	logging.Logger.Println("Database connection initialized:", service.Db)

	// 初始化 Gin 路由
	router := gin.Default()
	Router.InitRouter(router)

	// 启动 Gin 服务器
	if err := router.Run(); err != nil {
		logging.Logger.Fatalf("failed to run server: %v", err)
	}
}
