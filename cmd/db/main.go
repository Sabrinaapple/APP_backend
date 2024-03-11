package main

import (
	"freefrom.space/cms-service/internal/infra"
	"freefrom.space/cms-service/internal/model"
	"freefrom.space/cms-service/internal/model/schema"
	"gorm.io/gorm"
)

func main() {
	db := infra.New(model.NewEnv())
	createTable(db)
	return
}

func createTable(db *gorm.DB) {
	_ = db.AutoMigrate(
		&schema.User{},
	)
}
