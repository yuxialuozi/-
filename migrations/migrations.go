package migrations

import (
	"gorm.io/gorm"
	"simpledouyin/logging"
	"simpledouyin/model"
)

func AutoMigrate(DB *gorm.DB) {
	models := []interface{}{
		&model.Author{},
		&model.Video{},
		&model.UserLove{},
		&model.Comment{},
		&model.Relation{},
		&model.Message{},
	}

	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			logging.Logger.Fatalf("failed to auto-migrate model %v: %v", model, err)
		} else {
			logging.Logger.Infof("Successfully migrated model: %v", model)
		}
	}
}
