package migration

import (
	"goapi/config"
	"goapi/graph/model"
)

func MigrateTable() {
	db := config.GetDB()

	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{})
}
