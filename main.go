package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simpledouyin/Router"
	"simpledouyin/role"
	"simpledouyin/service"
)

func main() {
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:Lycdemima1@@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                 // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                               // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{}) //连接数据库

	if err != nil {
		fmt.Println(err)
	}

	//启动路由
	router := gin.Default()

	DB.AutoMigrate(&role.Author{})   //自动建表
	DB.AutoMigrate(&role.Video{})    //自动建表
	DB.AutoMigrate(&role.UserLove{}) //自动建表

	DB.AutoMigrate(&role.Comment{})  //自动建表
	DB.AutoMigrate(&role.Relation{}) //自动建表
	DB.AutoMigrate(&role.Message{})  //自动建表

	service.Db = DB
	fmt.Println(service.Db)

	//初始化路由，实现各种接口
	Router.InitRouter(router)

	router.Run() //设置运行接口，默认为8080
}
