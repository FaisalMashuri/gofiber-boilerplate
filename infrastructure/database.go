package infrastructure

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"gofiber-boilerplate/config"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.Username,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Database,
	)

	//define database logger
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dbConfig := gorm.Config{
		Logger: dbLogger,
	}

	db, err := gorm.Open(postgres.Open(dsn), &dbConfig)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
	// Todo : add struct you want to migrate
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
