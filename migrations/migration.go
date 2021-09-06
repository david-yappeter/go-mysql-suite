package migrations

import (
	"myapp/config"
	"myapp/entity"
)

func getModels() []interface{} {
	return []interface{}{&entity.User{}}
}

func MigrateTable() {
	db := config.GetDB()
	db.AutoMigrate(getModels()...)
}
