package db

import (
	"app/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Setup() {
	var err error
	Database, err = gorm.Open(sqlite.Open( config.Config.Database_Url ), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}