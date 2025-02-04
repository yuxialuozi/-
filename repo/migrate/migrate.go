package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"simpledouyin/config"
	"simpledouyin/repo/model"

	"gorm.io/gorm"
)

func main() {
	var err error
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN: config.EnvConfig.GetDSN(),
			}), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		panic(fmt.Errorf("db connection failed: %v", err))
	}
	err = db.AutoMigrate(&model.UserToken{}, &model.User{}, &model.Video{}, &model.Comment{}, &model.Relation{}, &model.Favorite{})
	if err != nil {
		panic(fmt.Errorf("db migrate failed: %v", err))
	}
}
