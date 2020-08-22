package repository

import (
	"github.com/jinzhu/gorm"

	"flight-api/app/core"
)

func getDB() *gorm.DB {
	return core.GetDB()
}
