package database

import (

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
	// Add other tables as needed for your application
}
