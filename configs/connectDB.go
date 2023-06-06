package configs

import (
	"notesApp/schemas"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DefaultDB *gorm.DB

func ConnectDB() {
	// connecting to DB
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// auto migrating tables
	db.AutoMigrate(&schemas.Notes{}, &schemas.User{})
	DefaultDB = db
}
