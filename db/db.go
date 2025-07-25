package db

import (
	"api/db/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDb() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=golang port=5433 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "api.",
		},
	})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Product{},
		&models.OrderItem{},
		&models.Category{},
		&models.Review{},
	)
	if err != nil {
		panic("Failed to migrate models")
	}

	return db
}

func CreateEntry(db *gorm.DB, model interface{}) error {
	result := db.Create(model)
	if result.Error != nil {
		return fmt.Errorf("failed to create entry: %v", result.Error)
	}
	return nil
}
