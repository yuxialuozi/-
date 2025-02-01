package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"simpledouyin/config"
	"simpledouyin/model2"
)

// Querier Dynamic SQL
type Querier interface {
	// FilterWithNameAndRole SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "dal",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	dsn := config.GetDBConfig()
	gormdb, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	g.UseDB(gormdb)

	// Generate basic type-safe DAO API for struct `model2.User` following conventions
	g.ApplyBasic(model2.UserToken{}, model2.Video{}, model2.User{}, model2.Comment{}, model2.Relation{}, model2.Favorite{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface
	g.ApplyInterface(func(Querier) {}, model2.UserToken{}, model2.Video{}, model2.User{}, model2.Comment{})

	// Generate the code
	g.Execute()
}
