package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/visaramadhan/project-dashboard-ecommerce-visa.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(makePostgresString(cfg)), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Database")
	}

	if cfg.DBMigrate {
		err = Migrate(db)
	}
	if err != nil {
		if err = Migrate(db); err != nil {
			return nil, fmt.Errorf("Failed to migrate Database")
		}

		if cfg.DBSeeding {
			err = SeedAll(gorm.DB{})
		}
		if err != nil {
			if err = SeedAll(gorm.DB{}); err != nil {
				return nil, err
			}
			return db, nil
		}
	}

	return db, nil
}

func makePostgresString(cfg config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
}
