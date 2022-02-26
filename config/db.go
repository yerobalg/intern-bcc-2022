package config

import (
	"clean-arch-2/app/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
	// "log"
	"os"
	// "time"
)

func InitDB() (*gorm.DB, error) {
	Init(".database")
	db, error := gorm.Open(
		postgres.Open(fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_DBNAME"),
		)),
		&gorm.Config{
		},
	)

	if error != nil {
		return nil, error
	}
	fmt.Println("Successfully connected to database!")

	db.AutoMigrate(
		// &models.Roles{},
		&models.Users{},
	)
	return db, error
}
